package app

import (
	"fmt"

	"github.com/pragmatically-dev/bookstore_users_api/logger"
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
	initializeRoutes() //initialize the routes
	logger.Info("About to start the application")
	router.Run()
}

//Middleware Json
func Json() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}
