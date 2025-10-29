package match

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sumoweb/server/internal/services/match"
)

func HandleEnqueueFetchMatches(c *gin.Context) {
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

	taskID, err := match.EnqueueFetchMatches(division, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":  "Match fetch task enqueued successfully",
		"task_id":  taskID,
		"division": division,
		"day":      day,
	})
}
