package domain

import "time"

type Score struct {
	HomeTeam int `json:"homeTeam"`
	AwayTeam int `json:"awayTeam"`
}

type Match struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	HomeTeam string    `json:"homeTeam"`
	AwayTeam string    `json:"awayTeam"`
	Score    Score     `json:"score"`
}

type Bet struct {
	ID      int   `json:"id"`
	IDMatch int   `json:"idMatch"`
	Score   Score `json:"score"`
}
