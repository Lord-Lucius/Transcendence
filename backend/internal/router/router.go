package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// Routes
	api := router.Group("/api")
	{
		// Test route
		api.GET("/posts", getPosts)
		api.POST("/posts", createPost)
	}
}

func getPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "TODO: Implement getPosts",
	})
}

func createPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "TODO: Implement createPost",
	})
}
