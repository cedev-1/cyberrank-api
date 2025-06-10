package handlers

import (
	"github.com/cedev-1/cyberrank-api/internal/models"
	"github.com/cedev-1/cyberrank-api/internal/scrapers"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTryHackMeRank(c *gin.Context) {
    username := c.Param("username")
    
    rank, err := scrapers.GetTryHackMeRank(username)
    if err != nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{
            Error: "User not found or error fetching data",
        })
        return
    }
    
    c.JSON(http.StatusOK, models.Profile{
        Username: username,
        Rank:     fmt.Sprintf("%d", rank),
        Platform: "tryhackme",
    })
}
