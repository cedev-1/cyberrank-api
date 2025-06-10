package handlers

import (
	"net/http"

	"github.com/cedev-1/cyberrank-api/internal/models"
	"github.com/cedev-1/cyberrank-api/internal/scrapers"

	"github.com/gin-gonic/gin"
)

func GetRootMeRank(c *gin.Context) {
	username := c.Param("username")

	detailed := c.Query("detailed") == "true"

	if detailed {
		profile, err := scrapers.GetRootMeProfile(username)
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "User not found or error fetching data",
			})
			return
		}

		c.JSON(http.StatusOK, models.Profile{
			Username:   username,
			Rank:       profile.Rank,
			Platform:   "root-me",
			Categories: profile.Categories,
			Overall:    profile.Overall,
		})
	} else {
		rank, err := scrapers.GetRootMeRank(username)
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "User not found or error fetching data",
			})
			return
		}

		c.JSON(http.StatusOK, models.Profile{
			Username: username,
			Rank:     rank,
			Platform: "root-me",
		})
	}
}
