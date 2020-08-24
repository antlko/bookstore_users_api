package app

import (
	"github.com/Anatol-e/bookstore_users_api/controllers/ping"
	"github.com/Anatol-e/bookstore_users_api/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUsers)
	router.POST("/users", users.CreateUser)
}