package rikishi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"sumoweb/server/internal/models"
	"sumoweb/server/internal/queue"
	"sumoweb/server/internal/workers/cache"
)

func HandleFetchRikishiProfile(ctx context.Context, t *asynq.Task) error {
	var p queue.FetchRikishiProfilePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	log.Printf("Processing task: type=%s rikishi_id=%s", t.Type(), p.RikishiID)

	enHTML, err := FetchProfile(p.RikishiID, models.English)
	if err != nil {
		return fmt.Errorf("failed to fetch EN profile: %w", err)
	}

	enProfile, err := ParseProfile(enHTML, models.English)
	if err != nil {
		return fmt.Errorf("failed to parse EN profile: %w", err)
	}

	enImageURL := ExtractImageURL(enHTML, models.English)
	if enImageURL != "" {
		imageServePath, err := DownloadProfileImage(p.RikishiID, models.English, enImageURL)
		if err == nil {
			enProfile.ImageURL = imageServePath
		} else {
			log.Printf("Warning: failed to download EN image for rikishi %s: %v", p.RikishiID, err)
		}
	}

	jpHTML, err := FetchProfile(p.RikishiID, models.Japanese)
	if err != nil {
		return fmt.Errorf("failed to fetch JP profile: %w", err)
	}

	jpProfile, err := ParseProfile(jpHTML, models.Japanese)
	if err != nil {
		return fmt.Errorf("failed to parse JP profile: %w", err)
	}

	jpImageURL := ExtractImageURL(jpHTML, models.Japanese)
	if jpImageURL != "" {
		imageServePath, err := DownloadProfileImage(p.RikishiID, models.Japanese, jpImageURL)
		if err == nil {
			jpProfile.ImageURL = imageServePath
		} else {
			log.Printf("Warning: failed to download JP image for rikishi %s: %v", p.RikishiID, err)
		}
	}

	rikishiProfile := &models.RikishiProfile{
		ID: p.RikishiID,
		Profiles: map[string]models.Profile{
			"en": enProfile,
			"jp": jpProfile,
		},
	}

	if err := cache.SaveRikishiProfileCache(rikishiProfile); err != nil {
		log.Printf("Warning: failed to save rikishi profile cache: %v", err)
	}

	log.Printf("Completed task: rikishi_id=%s", p.RikishiID)
	return nil
}
