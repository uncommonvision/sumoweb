package rikishi

import (
	"regexp"
	"strings"

	"sumoweb/server/internal/models"
)

func ParseProfile(html string, language models.Language) (models.Profile, error) {
	if language.ID == "en" {
		return parseEnglishProfile(html), nil
	}
	return parseJapaneseProfile(html), nil
}

func parseEnglishProfile(html string) models.Profile {
	profile := models.Profile{
		ImageURL: "",
	}

	profile.Stable = extractField(html, `<th[^>]*>\s*Stable\s*</th>\s*<td[^>]*>\s*(?:<a[^>]*>)?\s*([^<\n]+?)\s*(?:</a>)?\s*</td>`)
	profile.Name = extractField(html, `<th[^>]*>\s*Name\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.RingName = extractField(html, `<th[^>]*>\s*Ring Name\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.CurrentRank = extractField(html, `<th[^>]*>\s*Current Rank\s*</th>\s*<td[^>]*>(?:\s|<[^>]*>)*([^<\n]+?)(?:\s|<[^>]*>)*</td>`)
	profile.Birthday = extractField(html, `<th[^>]*>\s*Birthday\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.Birthplace = extractField(html, `<th[^>]*>\s*Birthplace\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.Height = extractField(html, `<th[^>]*>\s*Height\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.Weight = extractField(html, `<th[^>]*>\s*Weight\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.SignatureManeuver = extractField(html, `<th[^>]*>\s*Signature Maneuver\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)

	stableURLPattern := regexp.MustCompile(`href="(/EnSumoDataSumoBeya/[^"]+)"`)
	if matches := stableURLPattern.FindStringSubmatch(html); len(matches) > 1 {
		profile.StableURL = "https://www.sumo.or.jp" + matches[1]
	}

	return profile
}

func parseJapaneseProfile(html string) models.Profile {
	profile := models.Profile{
		ImageURL: "",
	}

	profile.Stable = extractField(html, `<th[^>]*>\s*所属部屋\s*</th>\s*<td[^>]*>\s*(?:<a[^>]*>)?\s*([^<\n]+?)\s*(?:</a>)?\s*</td>`)
	profile.Name = extractField(html, `<th[^>]*>\s*本名\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.RingName = extractField(html, `<th[^>]*>\s*しこ名履歴\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.CurrentRank = extractField(html, `([東西]?[^<]+枚目)`)
	profile.Birthday = extractField(html, `<th[^>]*>\s*生年月日\s*</th>\s*<td[^>]*>\s*([^<（\n]+)`)
	profile.Birthplace = extractField(html, `<th[^>]*>\s*出身地\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.Height = extractField(html, `<th[^>]*>\s*身長\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.Weight = extractField(html, `<th[^>]*>\s*体重\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)
	profile.SignatureManeuver = extractField(html, `<th[^>]*>\s*得意技\s*</th>\s*<td[^>]*>\s*([^<\n]+?)\s*</td>`)

	stableURLPattern := regexp.MustCompile(`href="(/ResultRikishiDataSumoBeya/[^"]+)"`)
	if matches := stableURLPattern.FindStringSubmatch(html); len(matches) > 1 {
		profile.StableURL = "https://www.sumo.or.jp" + matches[1]
	}

	return profile
}

func extractField(html, pattern string) string {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

func ExtractImageURL(html string, language models.Language) string {
	// Prioritize the specific sumo profile image format
	specificPattern := regexp.MustCompile(`<img[^>]+src="(/img/sumo_data/rikishi/270x474/[^"]+)"`)
	if matches := specificPattern.FindStringSubmatch(html); len(matches) > 1 {
		return "https://www.sumo.or.jp" + matches[1]
	}

	imagePattern := regexp.MustCompile(`<img[^>]+src="([^"]+)"[^>]*alt="[^"]*(?:profile|プロフィール|[^\s]+\s+[^\s]+)"`)
	if matches := imagePattern.FindStringSubmatch(html); len(matches) > 1 {
		imageURL := matches[1]
		if !strings.HasPrefix(imageURL, "http") {
			imageURL = "https://www.sumo.or.jp" + imageURL
		}
		return imageURL
	}

	fallbackPattern := regexp.MustCompile(`<img[^>]+src="(/[^"]+\.(?:jpg|png|jpeg))"`)
	if matches := fallbackPattern.FindStringSubmatch(html); len(matches) > 1 {
		return "https://www.sumo.or.jp" + matches[1]
	}

	return ""
}
