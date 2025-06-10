package main

import (
	"log"

	"github.com/cedev-1/cyberrank-api/internal/models/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    // CORS middleware
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
            return
        }
        c.Next()
    })
    
    api := r.Group("/api")
    {
        api.GET("/rootme/:username", handlers.GetRootMeRank)
        api.GET("/tryhackme/:username", handlers.GetTryHackMeRank)
    }
    
    log.Println("API server starting on :8080")
    r.Run(":8080")
}
