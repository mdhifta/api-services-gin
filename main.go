package main

import (
	"api-services/controllers"
	"api-services/initializers"
	"api-services/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()    // please your check .env
	initializers.Connection() // please check your connection database or comment if your not use db
}

func main() {
	router := gin.Default() //gin router initialization

	// route to web
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.Web)

	// route to api
	api := router.Group("/api")
	{
		// route without middleware
		api.POST("/sign-up", controllers.Register)
		api.POST("/sign-in", controllers.Login)

		// route with middleware
		api.Use(middleware.MiddlewareApi)
		api.GET("/", controllers.Services)

	}

	router.Run()
}
