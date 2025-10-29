package websocket

import (
	"encoding/json"
	"time"

	"sumoweb/server/internal/models"
)

const (
	MessageTypeConnectionAck = "CONNECTION_ACK"
	MessageTypePing          = "PING"
	MessageTypeUserIdentify  = "USER_IDENTIFY"
	MessageTypeUserJoined    = "USER_JOINED"
	MessageTypeUserLeft      = "USER_LEFT"
	MessageTypeChatMessage   = "CHAT_MESSAGE"
	MessageTypeMatchUpdate   = "MATCH_UPDATE"
)

type WSMessage struct {
	Type      string          `json:"type"`
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

func NewMessage(messageType string, data interface{}) (WSMessage, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return WSMessage{}, err
	}

	return WSMessage{
		Type:      messageType,
		Timestamp: time.Now(),
		Data:      jsonData,
	}, nil
}

type UserIdentifyPayload struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

type UserJoinedPayload struct {
	UserID    string `json:"userId"`
	UserName  string `json:"userName"`
	SessionID string `json:"sessionId"`
}

type UserLeftPayload struct {
	UserID    string `json:"userId"`
	UserName  string `json:"userName"`
	SessionID string `json:"sessionId"`
}

type MatchUpdatePayload struct {
	Division int            `json:"division"`
	Day      int            `json:"day"`
	Matches  []models.Match `json:"matches"` // Only changed matches (delta)
}
