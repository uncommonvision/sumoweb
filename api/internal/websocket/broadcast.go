package websocket

import (
	"fmt"
	"log"

	"sumoweb/server/internal/models"
)

func BroadcastMatchUpdates(division, day int, matches []models.Match) error {
	channelKey := fmt.Sprintf("sumo:division:%d:day:%d", division, day)

	err := Pool.BroadcastMatchUpdate(channelKey, division, day, matches)
	if err != nil {
		return fmt.Errorf("failed to broadcast: %w", err)
	}

	log.Printf("Broadcast MATCH_UPDATE to channel %s (%d matches)", channelKey, len(matches))
	return nil
}
