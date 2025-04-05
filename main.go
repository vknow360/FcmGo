package main

import (
	"os"

	"github.com/vknow360/FcmGo/config"
	"github.com/vknow360/FcmGo/firebase"
	"github.com/vknow360/FcmGo/routes"
	"github.com/vknow360/FcmGo/utils"
)

func main() {

	config.LoadEnv()

	serviceAccountJSON := config.GetEnv("SERVICE_ACCOUNT_JSON", "default-service-account-json")
	if serviceAccountJSON == "" {
		utils.ErrorLogger.Println("Failed to get service account json")
		os.Exit(1)
	}

	if err := firebase.InitFirebase([]byte(serviceAccountJSON)); err != nil {
		utils.ErrorLogger.Printf("Firebase init error: %v", err)
		os.Exit(1)
	}

	r := routes.RegisterRoutes()

	r.Run(":8080")
}
