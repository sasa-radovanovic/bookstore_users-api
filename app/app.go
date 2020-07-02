package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication starts the app
func StartApplication() {
	mapUrls()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
