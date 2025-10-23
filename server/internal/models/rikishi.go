package models

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	English  = Language{ID: "en", Name: "English"}
	Japanese = Language{ID: "jp", Name: "Japanese"}
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
