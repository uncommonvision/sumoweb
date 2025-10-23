package models

type Rikishi struct {
	ID             string             `json:"id"`
	ShikonaJP      string             `json:"shikona_jp"`
	ShikonaEN      string             `json:"shikona_en"`
	BanzukeNameJP  string             `json:"banzuke_name_jp"`
	BanzukeNameEN  string             `json:"banzuke_name_en"`
	Won            int                `json:"won"`
	Lost           int                `json:"lost"`
	RestJP         string             `json:"rest_jp"`
	RestEN         string             `json:"rest_en"`
	Result         bool               `json:"result"`
	KyokaiMemberID string             `json:"kyokai_member_id"`
	Profile        *RikishiProfile    `json:"profile,omitempty"`
}

type Match struct {
	East      Rikishi `json:"east"`
	West      Rikishi `json:"west"`
	JudgeID   string  `json:"judge_id"`
	TechnicID string  `json:"technic_id"`
}
