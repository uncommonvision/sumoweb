package models

import "time"

type MatchCache struct {
	Division  int       `json:"division"`
	Day       int       `json:"day"`
	FetchedAt time.Time `json:"fetched_at"`
	Matches   []Match   `json:"matches"`
}
