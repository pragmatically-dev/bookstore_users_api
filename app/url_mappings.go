package app

import (
	"github.com/pragmatically-dev/bookstore_users_api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
