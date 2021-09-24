package app

import "github.com/babadee08/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/user/:user_id", controllers.GetUser)
	router.POST("/user/search", controllers.SearchUser)
	router.POST("/user", controllers.CreateUser)
}
