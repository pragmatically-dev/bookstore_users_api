package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication this function start the whole app
func StartApplication() {
	fmt.Println("app running")
	mapUrls() //initialize the routes
	router.Run()
}
