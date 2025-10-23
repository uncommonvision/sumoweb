package rikishi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"sumoweb/server/internal/models"
)

func FetchProfile(rikishiID string, language models.Language) (string, error) {
	var url string

	if language.ID == "en" {
		url = fmt.Sprintf("https://www.sumo.or.jp/EnSumoDataRikishi/profile/%s/", rikishiID)
	} else {
		url = fmt.Sprintf("https://www.sumo.or.jp/ResultRikishiData/profile/%s/", rikishiID)
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch %s profile: %w", language.Name, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
