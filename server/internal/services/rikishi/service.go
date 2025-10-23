package rikishi

import (
	"fmt"

	"sumoweb/server/internal/models"
	"sumoweb/server/internal/workers/cache"
	rikishiWorker "sumoweb/server/internal/workers/rikishi"
)

func GetRikishiProfile(rikishiID string) (*models.RikishiProfile, error) {
	if rikishiID == "" {
		return nil, fmt.Errorf("rikishi ID is required")
	}

	cachedProfile, err := cache.GetRikishiProfileCache(rikishiID)
	if err == nil && cachedProfile != nil {
		return cachedProfile, nil
	}

	enChan := make(chan profileResult)
	jpChan := make(chan profileResult)

	go fetchAndParse(rikishiID, models.English, enChan)
	go fetchAndParse(rikishiID, models.Japanese, jpChan)

	enResult := <-enChan
	jpResult := <-jpChan

	if enResult.err != nil && jpResult.err != nil {
		return nil, fmt.Errorf("failed to fetch profiles: EN=%v, JP=%v", enResult.err, jpResult.err)
	}

	profiles := make(map[string]models.Profile)
	if enResult.err == nil {
		profiles["en"] = enResult.profile
	}
	if jpResult.err == nil {
		profiles["jp"] = jpResult.profile
	}

	rikishiProfile := &models.RikishiProfile{
		ID:       rikishiID,
		Profiles: profiles,
	}

	go cache.SaveRikishiProfileCache(rikishiProfile)

	return rikishiProfile, nil
}

type profileResult struct {
	profile models.Profile
	err     error
}

func fetchAndParse(id string, language models.Language, ch chan<- profileResult) {
	html, err := rikishiWorker.FetchProfile(id, language)
	if err != nil {
		ch <- profileResult{err: err}
		return
	}

	profile, err := rikishiWorker.ParseProfile(html, language)
	if err != nil {
		ch <- profileResult{err: err}
		return
	}

	imageURL := rikishiWorker.ExtractImageURL(html, language)
	if imageURL != "" {
		imageServePath, err := rikishiWorker.DownloadProfileImage(id, language, imageURL)
		if err == nil {
			profile.ImageURL = imageServePath
		}
	}

	ch <- profileResult{profile: profile, err: nil}
}
