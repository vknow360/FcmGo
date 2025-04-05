package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vknow360/FcmGo/utils"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		utils.ErrorLogger.Println(".env file not found or couldn't be loaded. Error: ", err.Error())
	}
}

func GetEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
