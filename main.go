package main

import (
	"go-url-shortner/config"
	"go-url-shortner/models"
	"go-url-shortner/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Url{})

	r := routes.SetupRouter()

	r.Run(":8080")
}