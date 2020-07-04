package app

import (
	ping "github.com/sasa-radovanovic/bookstore_users-api/controllers/ping_controller"
	"github.com/sasa-radovanovic/bookstore_users-api/controllers/users"
)

func mapUrls() {
	// Ping pong
	router.GET("/ping", ping.Ping)
	// Create user
	router.POST("/users", users.CreateUser)
	// Retrieve a user
	router.GET("/users/:user_id", users.GetUser)
	// Search users
	// router.GET("/users/search", controllers.SearchUser)
}
