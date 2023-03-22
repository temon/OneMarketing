package main

import (
	"log"
	appInit2 "oneMarketing/internal/app/appInit"
	model2 "oneMarketing/internal/app/model"
)

func init() {
	appInit2.LoadEnvVariables()
	appInit2.DBConnect()
}

func main() {
	errorMigration := appInit2.DB.AutoMigrate(
		&model2.User{},
		&model2.SubscriptionPlan{},
	)

	if errorMigration != nil {
		log.Fatal("DB migration Up error")
	}
}
