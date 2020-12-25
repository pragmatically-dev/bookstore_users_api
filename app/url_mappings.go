package app

import (
	"github.com/pragmatically-dev/bookstore_users_api/controllers/ping"
	"github.com/pragmatically-dev/bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:id", user.GetUser)
	//	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", user.CreateUser)
}
