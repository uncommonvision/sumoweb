package websocket

import (
	"log"

	ws "sumoweb/server/internal/websocket"
)

func HandlePing(conn *ws.Connection, msg ws.WSMessage) {
	log.Printf("PING received from connection %s (uuid: %s)", conn.ID, conn.UUID)
}
