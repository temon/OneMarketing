package appInit

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

func LoadEnvVariables() {
	projectDirName := "oneMarketing"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/internal/app/.env`)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
