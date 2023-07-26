package main

import (
	"api-services/initializers"
	"api-services/models"
)

func init() {
	initializers.LoadEnv()
	initializers.Connection()
}

func main() {
	initializers.DB.AutoMigrate(&models.Users{})
}
