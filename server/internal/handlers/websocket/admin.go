package websocket

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ws "sumoweb/server/internal/websocket"
)

func HandleBroadcastMatches(c *gin.Context) {
	division, err := strconv.Atoi(c.Param("division"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid division parameter"})
		return
	}

	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid day parameter"})
		return
	}

	if err := ws.BroadcastMatchUpdates(division, day); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Match updates broadcast successfully",
		"division": division,
		"day":      day,
	})
}
