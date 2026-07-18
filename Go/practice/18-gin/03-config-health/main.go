package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port  string
	Ready bool
}

func loadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Port:  port,
		Ready: os.Getenv("READY") == "true",
	}
}

func main() {
	cfg := loadConfig()

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/readyz", func(c *gin.Context) {
		if !cfg.Ready {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not_ready"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	_ = router.Run(":" + cfg.Port)
}
