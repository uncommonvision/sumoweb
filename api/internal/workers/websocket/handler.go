package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"sumoweb/server/internal/queue"
	matchService "sumoweb/server/internal/services/match"
	"sumoweb/server/internal/websocket"
)

func HandleBroadcastMatches(ctx context.Context, t *asynq.Task) error {
	var p queue.BroadcastMatchesPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	log.Printf("Processing broadcast task: division=%d day=%d with %d changed matches", p.Division, p.Day, len(p.Matches))

	if len(p.Matches) == 0 {
		log.Printf("No matches to broadcast for division=%d day=%d", p.Division, p.Day)
		return nil
	}

	enrichedMatches := matchService.EnrichMatchesWithCachedProfiles(p.Matches, p.Division, p.Day)

	if err := websocket.BroadcastMatchUpdates(p.Division, p.Day, enrichedMatches); err != nil {
		return fmt.Errorf("failed to broadcast: %w", err)
	}

	log.Printf("Completed broadcast: division=%d day=%d (%d changed matches)", p.Division, p.Day, len(enrichedMatches))
	return nil
}
