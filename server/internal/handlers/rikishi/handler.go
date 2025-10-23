package rikishi

import (
	"fmt"
	"net/http"
	"path/filepath"

	"sumoweb/server/internal/services/rikishi"
	"github.com/gin-gonic/gin"
)

func HandleGetRikishi(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rikishi ID is required"})
		return
	}

	profile, err := rikishi.GetRikishiProfile(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func HandleGetRikishiImage(c *gin.Context) {
	id := c.Param("id")
	lang := c.Param("lang")

	if id == "" || lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rikishi ID and language are required"})
		return
	}

	if lang != "en" && lang != "jp" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "language must be 'en' or 'jp'"})
		return
	}

	filename := fmt.Sprintf("%s-%s.png", id, lang)
	imagePath := filepath.Join("public", "rikishi", "profile", filename)

	c.File(imagePath)
}
