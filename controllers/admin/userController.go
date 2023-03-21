package admin

import (
	"github.com/gin-gonic/gin"
	"oneMarketing/appInit"
	"oneMarketing/models"
)

type UserRequest struct {
	Email    string
	Name     string
	Password string
}

// Admin API

func UserCreate(c *gin.Context) {

	var userReqData UserRequest
	parseError := c.BindJSON(&userReqData)

	if parseError != nil {
		c.Status(400)
		return
	}

	user := models.User{
		Email:    userReqData.Email,
		Name:     userReqData.Name,
		Password: userReqData.Password,
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

func UserUpdate(c *gin.Context) {
	id := c.Param("id")

	var userReqData UserRequest
	parseError := c.BindJSON(&userReqData)

	if parseError != nil {
		c.Status(400)
		return
	}

	var user models.User
	appInit.DB.First(&user, id)

	appInit.DB.Model(&user).Updates(&models.User{
		Email:    userReqData.Email,
		Name:     userReqData.Name,
		Password: userReqData.Password,
	})

	c.JSON(200, gin.H{
		"users": user,
	})
}

func UserIndex(c *gin.Context) {
	var users []models.User
	appInit.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserShow(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	appInit.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	appInit.DB.Delete(&models.User{}, id)

	c.Status(200)
}
