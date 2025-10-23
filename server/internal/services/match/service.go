package match

import (
	"fmt"
	"sync"

	"sumoweb/server/internal/models"
	rikishiService "sumoweb/server/internal/services/rikishi"
	"sumoweb/server/internal/workers/cache"
	matchWorker "sumoweb/server/internal/workers/match"
)

func GetMatches(division, day int) ([]models.Match, error) {
	if err := validateParams(division, day); err != nil {
		return nil, err
	}

	cachedData, err := cache.GetMatchCache(division, day)
	if err == nil && cachedData != nil {
		return cachedData.Matches, nil
	}

	response, err := matchWorker.FetchTorikumiData(division, day)
	if err != nil {
		return nil, err
	}

	matches := matchWorker.ParseMatches(response)

	enrichedMatches, err := enrichMatchesWithProfiles(matches)
	if err != nil {
		enrichedMatches = matches
	}

	go cache.SaveMatchCache(division, day, enrichedMatches)

	return enrichedMatches, nil
}

func enrichMatchesWithProfiles(matches []models.Match) ([]models.Match, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	enrichedMatches := make([]models.Match, len(matches))

	for i, match := range matches {
		enrichedMatches[i] = match

		wg.Add(2)

		go func(idx int, rikishiID string) {
			defer wg.Done()
			profile, err := rikishiService.GetRikishiProfile(rikishiID)
			if err == nil {
				mu.Lock()
				enrichedMatches[idx].East.Profile = profile
				mu.Unlock()
			}
		}(i, match.East.ID)

		go func(idx int, rikishiID string) {
			defer wg.Done()
			profile, err := rikishiService.GetRikishiProfile(rikishiID)
			if err == nil {
				mu.Lock()
				enrichedMatches[idx].West.Profile = profile
				mu.Unlock()
			}
		}(i, match.West.ID)
	}

	wg.Wait()
	return enrichedMatches, nil
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
