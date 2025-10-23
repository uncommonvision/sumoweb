package rikishi

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"sumoweb/server/internal/models"
)

func DownloadProfileImage(rikishiID string, language models.Language, imageURL string) (string, error) {
	if imageURL == "" {
		return "", fmt.Errorf("no image URL provided")
	}

	filename := fmt.Sprintf("%s-%s.png", rikishiID, language.ID)
	filePath := filepath.Join("public", "rikishi", "profile", filename)

	if _, err := os.Stat(filePath); err == nil {
		return fmt.Sprintf("/api/v1/rikishi/%s/%s/image.png", rikishiID, language.ID), nil
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	resp, err := http.Get(imageURL)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write image: %w", err)
	}

	return fmt.Sprintf("/api/v1/rikishi/%s/%s/image.png", rikishiID, language.ID), nil
}
