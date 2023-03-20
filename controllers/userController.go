package controllers

import (
	"github.com/gin-gonic/gin"
	"oneMarketing/appInit"
	"oneMarketing/models"
)

func UserCreate(c *gin.Context) {

	user := models.User{
		Email:    "eko@testing.com",
		Name:     "eko",
		Password: "uuuu",
	}

	result := appInit.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": user,
	})
}
