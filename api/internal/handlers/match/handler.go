package match

import (
	"net/http"
	"strconv"

	"sumoweb/server/internal/services/match"
	"github.com/gin-gonic/gin"
)

func HandleGetMatches(c *gin.Context) {
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

	matches, err := match.GetMatches(division, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}
