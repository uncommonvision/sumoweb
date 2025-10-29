package rikishi

import (
	"context"
	"fmt"
	"log"

	"sumoweb/server/internal/models"
	"sumoweb/server/internal/queue"
	"sumoweb/server/internal/workers/cache"
)

func GetRikishiProfile(rikishiID string) (*models.RikishiProfile, error) {
	if rikishiID == "" {
		return nil, fmt.Errorf("rikishi ID is required")
	}

	cachedProfile, err := cache.GetRikishiProfileCache(rikishiID)
	if err == nil && cachedProfile != nil {
		return cachedProfile, nil
	}

	return nil, fmt.Errorf("rikishi profile not cached: %s", rikishiID)
}

func DiscoverMissingProfiles() ([]string, error) {
	matchCaches, err := cache.ListAllMatchCaches()
	if err != nil {
		return nil, fmt.Errorf("failed to list match caches: %w", err)
	}

	log.Printf("Found %d match cache files", len(matchCaches))

	rikishiIDs := make(map[string]bool)

	for _, matchCache := range matchCaches {
		for _, match := range matchCache.Matches {
			rikishiIDs[match.East.ID] = true
			rikishiIDs[match.West.ID] = true
		}
	}

	log.Printf("Found %d unique rikishi IDs from matches", len(rikishiIDs))

	missingIDs := []string{}
	for id := range rikishiIDs {
		profile, err := cache.GetRikishiProfileCache(id)
		if err != nil || profile == nil {
			missingIDs = append(missingIDs, id)
		}
	}

	log.Printf("Found %d rikishi with missing profiles", len(missingIDs))

	return missingIDs, nil
}

func EnqueueMissingProfiles() (int, error) {
	missingIDs, err := DiscoverMissingProfiles()
	if err != nil {
		return 0, fmt.Errorf("failed to discover missing profiles: %w", err)
	}

	producer := queue.GetProducer()
	if producer == nil {
		return 0, fmt.Errorf("queue not available")
	}

	enqueued := 0
	for _, id := range missingIDs {
		_, err := producer.EnqueueFetchRikishiProfile(context.Background(), id)
		if err != nil {
			log.Printf("Failed to enqueue profile fetch for rikishi %s: %v", id, err)
			continue
		}
		enqueued++
	}

	log.Printf("Enqueued %d rikishi profile fetch tasks", enqueued)
	return enqueued, nil
}
