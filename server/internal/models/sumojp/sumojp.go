package sumojp

type Response struct {
	TorHtml      string  `json:"TorHtml"`
	FinHtml      string  `json:"FinHtml"`
	Result       string  `json:"Result"`
	DayName      string  `json:"dayName"`
	DayHead      string  `json:"dayHead"`
	DayStr       string  `json:"dayStr"`
	KakuName     string  `json:"kakuName"`
	PdfUrl       string  `json:"PdfUrl"`
	En           string  `json:"en"`
	BashoID      int     `json:"basho_id"`
	KakuzukeID   string  `json:"kakuzuke_id"`
	Day          string  `json:"day"`
	DispFlg      int     `json:"dispFlg"`
	TorikumiData []Match `json:"TorikumiData"`
}

type Match struct {
	Judge          int     `json:"judge"`
	TechnicName    string  `json:"technic_name"`
	TechnicNameEng string  `json:"technic_name_eng"`
	TechnicID      int     `json:"technic_id"`
	East           Rikishi `json:"east"`
	West           Rikishi `json:"west"`
}

type Rikishi struct {
	RikishiID      int    `json:"rikishi_id"`
	KakuID         string `json:"kaku_id"`
	BanzukeName    string `json:"banzuke_name"`
	BanzukeNameEng string `json:"banzuke_name_eng"`
	Shikona        string `json:"shikona"`
	ShikonaEng     string `json:"shikona_eng"`
	WonNumber      int    `json:"won_number"`
	LostNumber     int    `json:"lost_number"`
	Rest           string `json:"rest"`
	RestEng        string `json:"rest_eng"`
	ResultImage    string `json:"result_image"`
	Alt            string `json:"alt"`
	KyokaiMemberID string `json:"kyokai_member_id"`
	WinLose        string `json:"win_lose"`
}
