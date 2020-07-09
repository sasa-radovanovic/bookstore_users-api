package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sasa-radovanovic/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

// StartApplication starts the app
func StartApplication() {
	mapUrls()

	logger.Info("About to start an application")
	router.Run(":8080")
}
