package match

import (
	"fmt"
	"strconv"

	"sumoweb/server/internal/models"
	"sumoweb/server/internal/models/sumojp"
)

func ParseMatches(response *sumojp.Response) []models.Match {
	matches := make([]models.Match, 0, len(response.TorikumiData))

	for _, rawMatch := range response.TorikumiData {
		match := models.Match{
			East:      parseRikishi(rawMatch.East),
			West:      parseRikishi(rawMatch.West),
			JudgeID:   strconv.Itoa(rawMatch.Judge),
			TechnicID: strconv.Itoa(rawMatch.TechnicID),
		}
		matches = append(matches, match)
	}

	return matches
}

func parseRikishi(raw sumojp.Rikishi) models.Rikishi {
	return models.Rikishi{
		ID:             fmt.Sprintf("%d", raw.RikishiID),
		ShikonaJP:      raw.Shikona,
		ShikonaEN:      raw.ShikonaEng,
		BanzukeNameJP:  raw.BanzukeName,
		BanzukeNameEN:  raw.BanzukeNameEng,
		Won:            raw.WonNumber,
		Lost:           raw.LostNumber,
		RestJP:         raw.Rest,
		RestEN:         raw.RestEng,
		Result:         raw.ResultImage == "result_ic01.gif",
		KyokaiMemberID: raw.KyokaiMemberID,
	}
}
