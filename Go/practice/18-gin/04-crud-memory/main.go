package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	jobs := map[string]string{}
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		jobs[id] = "queued"
		c.JSON(http.StatusCreated, gin.H{"id": id, "status": jobs[id]})
	})

	router.GET("/jobs", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"jobs": jobs})
	})

	router.DELETE("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		delete(jobs, id)
		c.JSON(http.StatusOK, gin.H{"deleted": id})
	})

	_ = router.Run(":8080")
}
