package models

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	English  = Language{ID: "en", Name: "English"}
	Japanese = Language{ID: "jp", Name: "Japanese"}
)

const (
	MatchResultPending = -1
	MatchResultLost    = 0
	MatchResultWon     = 1
)

type Profile struct {
	Stable            string `json:"stable"`
	StableURL         string `json:"stable_url"`
	Name              string `json:"name"`
	RingName          string `json:"ring_name"`
	CurrentRank       string `json:"current_rank"`
	Birthday          string `json:"birthday"`
	Birthplace        string `json:"birthplace"`
	Height            string `json:"height"`
	Weight            string `json:"weight"`
	SignatureManeuver string `json:"signature_maneuver"`
	ImageURL          string `json:"image_url"`
}

type RikishiProfile struct {
	ID       string             `json:"id"`
	Profiles map[string]Profile `json:"profiles"`
}

type Rikishi struct {
	ID             string          `json:"id"`
	ShikonaJP      string          `json:"shikona_jp"`
	ShikonaEN      string          `json:"shikona_en"`
	BanzukeNameJP  string          `json:"banzuke_name_jp"`
	BanzukeNameEN  string          `json:"banzuke_name_en"`
	Won            int             `json:"won"`
	Lost           int             `json:"lost"`
	RestJP         string          `json:"rest_jp"`
	RestEN         string          `json:"rest_en"`
	Result         int             `json:"result"`
	KyokaiMemberID string          `json:"kyokai_member_id"`
	Profile        *RikishiProfile `json:"profile,omitempty"`
}

type Match struct {
	East      Rikishi `json:"east"`
	West      Rikishi `json:"west"`
	JudgeID   string  `json:"judge_id"`
	TechnicID string  `json:"technic_id"`
}
