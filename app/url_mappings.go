package app

import (
	"github.com/babadee08/bookstore_users-api/controllers/ping"
	"github.com/babadee08/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", users.GetUser)
	router.POST("/user", users.CreateUser)
}
