package app

import (
	"fmt"

	"github.com/pragmatically-dev/bookstore_users_api/utils/config"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication this function start the whole app
func StartApplication() {
	config.Load()
	fmt.Println("app running")
	router.Use(Json())
	mapUrls() //initialize the routes
	router.Run()
}

//Middleware Json
func Json() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}
