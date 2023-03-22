package main

import (
	"github.com/gin-gonic/gin"
	"log"
	appInit2 "oneMarketing/internal/app/appInit"
	"oneMarketing/internal/app/controller/admin"
	"oneMarketing/internal/app/controller/apps"
	"oneMarketing/internal/app/middleware"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// admin
	r.POST("/api/v1/user", admin.UserCreate)
	r.PUT("/api/v1/user/:id", admin.UserUpdate)
	r.GET("/api/v1/user", admin.UserIndex)
	r.GET("/api/v1/user/:id", admin.UserShow)
	r.DELETE("/api/v1/user/:id", admin.UserDelete)

	// application
	r.POST("/api/v1/signUp", apps.Signup)
	r.POST("/api/v1/login", apps.Login)
	r.GET("/api/v1/me", middleware.JwtAuth, apps.Me)

	return r
}

func init() {
	appInit2.LoadEnvVariables()
	appInit2.DBConnect()
}

func main() {
	r := setupRouter()

	ginRunError := r.Run()

	if ginRunError != nil {
		log.Fatal("Can not run Gin")
	}
}
