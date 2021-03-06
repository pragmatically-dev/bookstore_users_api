package app

import (
	"github.com/pragmatically-dev/bookstore_users_api/controllers/ping"
	"github.com/pragmatically-dev/bookstore_users_api/controllers/user"
)

func initializeRoutes() {
	router.GET("/ping", ping.Ping)

	//users routes
	router.POST("/users", user.Create)
	router.GET("/users/:id", user.Get)
	router.PUT("/users/:id", user.Update)
	router.PATCH("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)

	router.GET("/internal/users/search", user.Search)
	router.POST("/users/login", user.Login)
}
