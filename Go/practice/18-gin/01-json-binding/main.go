package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateJobRequest struct {
	Name   string `json:"name" binding:"required"`
	Target string `json:"target" binding:"required"`
}

type CreateJobResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/jobs", func(c *gin.Context) {
		var request CreateJobRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		response := CreateJobResponse{
			ID:     "job-123",
			Status: "queued",
		}
		c.JSON(http.StatusCreated, response)
	})

	_ = router.Run(":8080")
}
