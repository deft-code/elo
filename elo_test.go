package elo

import "testing"

func TestCorrectScores(t *testing.T) {
	if LOSE != 0.0 {
		t.Errorf("LOSE(%v) != 0.0", LOSE)
	}
	if DRAW != 0.5 {
		t.Errorf("DRAW(%v) != 0.5", DRAW)
	}
	if WIN != 1.0 {
		t.Errorf("WIN(%v) != 1.0", WIN)
	}
}

func TestWikipediaExample1(t *testing.T) {
	rating := 1613
	series := []Match{{1609, LOSE}, {1477, DRAW}, {1388, WIN}, {1586, WIN}, {1720, LOSE}}
	new_rating := Rate(rating, series, UscfKFactor)
	if new_rating != 1601 {
		t.Errorf("%v != 1601", new_rating)
	}
}

func TestWikipediaExample2(t *testing.T) {
  rating := 1613
  series := []Match{{1609, LOSE}, {1477, DRAW}, {1388, WIN}, {1586, WIN}, {1720, DRAW}}
  new_rating := Rate(rating, series, UscfKFactor)
  if new_rating != 1617 {
    t.Errorf("%v != 1617", new_rating)
  }
}
