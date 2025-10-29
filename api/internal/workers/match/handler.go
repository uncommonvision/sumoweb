package match

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

func HandleFetchMatches(ctx context.Context, t *asynq.Task) error {
	var p queue.FetchMatchesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	log.Printf("Processing task: type=%s division=%d day=%d", t.Type(), p.Division, p.Day)

	response, err := FetchTorikumiData(p.Division, p.Day)
	if err != nil {
		return fmt.Errorf("failed to fetch torikumi data: %w", err)
	}

	matches := ParseMatches(response)

	oldCache, err := cache.GetMatchCache(p.Division, p.Day)
	if err != nil {
		log.Printf("Warning: failed to load old cache: %v", err)
	}

	var oldMatches []models.Match
	if oldCache != nil {
		oldMatches = oldCache.Matches
	}

	changedMatches := findChangedMatches(oldMatches, matches)

	if err := cache.SaveMatchCache(p.Division, p.Day, matches); err != nil {
		log.Printf("Warning: failed to save match cache: %v", err)
	}

	if len(changedMatches) > 0 {
		producer := queue.GetProducer()
		if producer != nil {
			if _, err := producer.EnqueueBroadcastMatches(context.Background(), p.Division, p.Day, changedMatches); err != nil {
				log.Printf("Warning: failed to enqueue broadcast task: %v", err)
			} else {
				log.Printf("Enqueued broadcast task for division=%d day=%d with %d changed matches", p.Division, p.Day, len(changedMatches))
			}
		}
	} else {
		log.Printf("No changes detected for division=%d day=%d, skipping broadcast", p.Division, p.Day)
	}

	log.Printf("Completed task: division=%d day=%d matches=%d", p.Division, p.Day, len(matches))
	return nil
}

func findChangedMatches(oldMatches, newMatches []models.Match) []models.Match {
	if oldMatches == nil {
		return newMatches
	}

	oldMap := make(map[string]models.Match)
	for _, m := range oldMatches {
		oldMap[m.Key()] = m
	}

	var changed []models.Match
	for _, newMatch := range newMatches {
		if oldMatch, exists := oldMap[newMatch.Key()]; exists {
			if newMatch.HasChangedFrom(oldMatch) {
				changed = append(changed, newMatch)
			}
		} else {
			changed = append(changed, newMatch)
		}
	}

	return changed
}
