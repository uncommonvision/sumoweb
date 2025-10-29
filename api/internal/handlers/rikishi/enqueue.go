package rikishi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sumoweb/server/internal/services/rikishi"
)

func HandleEnqueueRikishiProfiles(c *gin.Context) {
	count, err := rikishi.EnqueueMissingProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Rikishi profile fetch tasks enqueued successfully",
		"count":   count,
	})
}
