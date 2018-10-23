package models

import (
	"time"
)

// Player holds fields related to a player
type Player struct {
	ID            int     `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	OverallRating float64 `db:"overall_rating" json:"overall_rating"`
	OverallWins   int     `db:"overall_wins" json:"overall_wins"`
	OverallLosses int     `db:"overall_losses" json:"overall_losses"`
}

// Season holds fields related to a season
type Season struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	StartDate time.Time `db:"start_date" json:"start_date"`
	EndDate   time.Time `db:"end_date" json:"end_date"`
}

// SeasonalStats holds fields related to a players stats in a particular season
type SeasonalStats struct {
	ID       int     `db:"id" json:"id"`
	PlayerID int     `db:"player_id" json:"player_id"`
	SeasonID int     `db:"season_id" json:"season_id"`
	Rating   float64 `db:"rating" json:"rating"`
	Wins     int     `db:"wins" json:"wins"`
	Losses   int     `db:"losses" json:"losses"`
}

// Match holds fields relating to a match
type Match struct {
	ID       int `db:"id" json:"id"`
	SeasonID int `db:"season_id" json:"season_id"`
}

// MatchCompetitor holds fields related to a match played by a player and its affect on the player
type MatchCompetitor struct {
	ID             int     `db:"id" json:"id"`
	MatchID        int     `db:"match_id" json:"match_id"`
	PlayerID       int     `db:"player_id" json:"player_id"`
	StartingRating float64 `db:"starting_rating" json:"starting_rating"`
	RatingChange   float64 `db:"rating_change" json:"rating_change"`
	Winner         bool    `db:"winner" json:"winner"`
}
