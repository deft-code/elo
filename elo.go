package elo

import "math"

type ScoreType float64

const (
	LOSE ScoreType = iota * 0.5
	DRAW
	WIN
)

type KFactor func(rating int) float64

type Match struct {
	Rating int
	Score  ScoreType
}

func NewConstKFactor(k_factor float64) KFactor {
	return func(int) float64 { return k_factor }
}

func UscfKFactor(rating int) float64 {
	switch {
	case rating < 2100:
		return 32
	case rating < 2400:
		return 24
	default:
		return 16 // rating >= 2400
	}
	panic("Never get here!")
}

func expect(rating, other_rating int) float64 {
	return 1.0 / (1.0 + math.Pow(10, float64(other_rating-rating)/400.0))
}

func delta(rating int, series []Match) float64 {
	expected := 0.0
	actual := 0.0
	for _, res := range series {
		expected += expect(rating, res.Rating)
		actual += float64(res.Score)
	}
	return actual - expected
}

func Rate(rating int, series []Match, k_factor KFactor) int {
	return rating + int(math.Floor(0.5+k_factor(rating)*delta(rating, series)))
}

func RateWin(rating1, rating2 int, k_factor KFactor) (int, int) {
	one := Rate(rating1, []Match{{rating2, WIN}}, k_factor)
	two := Rate(rating2, []Match{{rating1, LOSE}}, k_factor)
	return one, two
}

func RateLose(rating1, rating2 int, k_factor KFactor) (int, int) {
	two, one := RateWin(rating2, rating1, k_factor)
	return one, two
}

func RateDraw(rating1, rating2 int, k_factor KFactor) (int, int) {
	one := Rate(rating1, []Match{{rating2, DRAW}}, k_factor)
	two := Rate(rating2, []Match{{rating1, DRAW}}, k_factor)
	return one, two
}
