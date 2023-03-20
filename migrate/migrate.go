package main

import (
	"log"
	"oneMarketing/appInit"
	"oneMarketing/models"
)

func init() {
	appInit.LoadEnvVariables()
	appInit.DBConnect()
}

func main() {
	errorMigration := appInit.DB.AutoMigrate(
		&models.User{},
		&models.SubscriptionPlan{},
	)

	if errorMigration != nil {
		log.Fatal("DB migration error")
	}
}
