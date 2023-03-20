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
	errorMigration := appInit.DB.Migrator().DropTable(
		&models.User{},
		&models.SubscriptionPlan{},
	)

	if errorMigration != nil {
		log.Fatal("DB migration Down error")
	}
}
