package main

import (
	"log"
	"oneMarketing/appInit"
	"oneMarketing/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/api/v1/user", controllers.UserCreate)

	return r
}

func init() {
	appInit.LoadEnvVariables()
	appInit.DBConnect()
}

func main() {
	r := setupRouter()

	ginRunError := r.Run()

	if ginRunError != nil {
		log.Fatal("Can not run Gin")
	}
}
