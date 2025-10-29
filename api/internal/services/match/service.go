package match

import (
	"context"
	"fmt"
	"log"

	"sumoweb/server/internal/models"
	"sumoweb/server/internal/queue"
	"sumoweb/server/internal/workers/cache"
)

func GetMatches(division, day int) ([]models.Match, error) {
	if err := validateParams(division, day); err != nil {
		return nil, err
	}

	cachedData, err := cache.GetMatchCache(division, day)
	if err != nil {
		return nil, fmt.Errorf("cache error: %w", err)
	}

	if cachedData == nil {
		return nil, fmt.Errorf("no cached data available for division %d, day %d", division, day)
	}

	enrichedMatches := EnrichMatchesWithCachedProfiles(cachedData.Matches, division, day)

	return enrichedMatches, nil
}

func EnrichMatchesWithCachedProfiles(matches []models.Match, division, day int) []models.Match {
	enriched := make([]models.Match, len(matches))

	for i, match := range matches {
		enriched[i] = match

		if profile, err := cache.GetRikishiProfileCache(match.East.ID); err == nil && profile != nil {
			enriched[i].East.Profiles = profile.Profiles
		} else {
			log.Printf("Profile not cached for rikishi ID %s (east, division: %d, day: %d)",
				match.East.ID, division, day)
		}

		if profile, err := cache.GetRikishiProfileCache(match.West.ID); err == nil && profile != nil {
			enriched[i].West.Profiles = profile.Profiles
		} else {
			log.Printf("Profile not cached for rikishi ID %s (west, division: %d, day: %d)",
				match.West.ID, division, day)
		}
	}

	return enriched
}

func EnqueueFetchMatches(division, day int) (string, error) {
	if err := validateParams(division, day); err != nil {
		return "", err
	}

	producer := queue.GetProducer()
	if producer == nil {
		return "", fmt.Errorf("queue not available")
	}

	taskID, err := producer.EnqueueFetchMatches(context.Background(), division, day)
	if err != nil {
		return "", fmt.Errorf("failed to enqueue fetch task: %w", err)
	}

	return taskID, nil
}

func validateParams(division, day int) error {
	if division < 1 || division > 6 {
		return fmt.Errorf("invalid division: %d (must be between 1 and 6)", division)
	}
	if day < 1 || day > 15 {
		return fmt.Errorf("invalid day: %d (must be between 1 and 15)", day)
	}
	return nil
}
