package queue

import "sumoweb/server/internal/models"

const (
	TypeFetchMatches        = "fetch:matches"
	TypeFetchRikishiProfile = "fetch:rikishi_profile"
	TypeBroadcastMatches    = "broadcast:matches"
)

type FetchMatchesPayload struct {
	Day      int `json:"day"`
	Division int `json:"division"`
}

type FetchRikishiProfilePayload struct {
	RikishiID string `json:"rikishi_id"`
}

type BroadcastMatchesPayload struct {
	Division int            `json:"division"`
	Day      int            `json:"day"`
	Matches  []models.Match `json:"matches"`
}
