package elo

import "math"

const kFactor = 20

// CalculateUpdatedRatings takes the rating of a winner and a loser and calculates their updated ratings.
func CalculateUpdatedRatings(winnerRating, loserRating float64) (winnerUpdatedRating, loserUpdatedRating float64) {
	winnerExpected, loserExpected := expectedRatings(winnerRating, loserRating)
	winnerUpdatedRating = winnerRating + kFactor*(1-winnerExpected)
	loserUpdatedRating = loserRating + kFactor*(-loserExpected)
	return winnerUpdatedRating, loserUpdatedRating
}

func transformedRatings(rating1, rating2 float64) (transformed1, transfromed2 float64) {
	return math.Pow(10, rating1/400), math.Pow(10, rating2/400)
}

func expectedRatings(rating1, rating2 float64) (expected1, expected2 float64) {
	transformed1, transformed2 := transformedRatings(rating1, rating2)
	denom := transformed1 + transformed2
	return transformed1 / denom, transformed2 / denom
}
