package websocket

import (
	"fmt"
	"log"

	"sumoweb/server/internal/models"
	"sumoweb/server/internal/services/match"
)

func BroadcastMatchUpdates(division, day int) error {
	matches, err := match.GetMatches(division, day)
	if err != nil {
		return fmt.Errorf("failed to fetch matches: %w", err)
	}

	channelKey := fmt.Sprintf("sumo:division:%d:day:%d", division, day)

	match := matches[len(matches)-1]

	match.East.Result = 0
	match.West.Result = 1

	err = Pool.BroadcastMatchUpdate(channelKey, division, day, []models.Match{match})
	if err != nil {
		return fmt.Errorf("failed to broadcast: %w", err)
	}

	log.Printf("Broadcast MATCH_UPDATE to channel %s (%d matches)", channelKey, len(matches))
	return nil
}
