package handlers

import (
	"net/http"

	"github.com/cedev-1/cyberrank-api/internal/scrapers"
	"github.com/gin-gonic/gin"
)

func GetTryHackMeRank(c *gin.Context) {
	username := c.Param("username")
	detailed := c.Query("detailed") == "true"

	profile, err := scrapers.GetTryHackMeRank(username, detailed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
