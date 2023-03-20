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
	r.PUT("/api/v1/user/:id", controllers.UserUpdate)
	r.GET("/api/v1/user", controllers.UserIndex)
	r.GET("/api/v1/user/:id", controllers.UserShow)
	r.DELETE("/api/v1/user/:id", controllers.UserDelete)

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
