package elo_test

import (
	"testing"

	"github.com/marquessv/slack-league/elo"
)

func TestCalculateUpdateTotal(t *testing.T) {
	winnerUpdatedRating, loserUpdatedRating := elo.CalculateUpdatedRatings(1400, 1400)
	if winnerUpdatedRating != 1410 {
		t.Errorf("Winners updated rating was incorrect, got: %f, expected: 1410", winnerUpdatedRating)
	}
	if loserUpdatedRating != 1390 {
		t.Errorf("Losers updated rating was incorrect, got: %f, expected 1390", loserUpdatedRating)
	}
}
