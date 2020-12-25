package app

import (
	"github.com/pragmatically-dev/bookstore_users_api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
}
