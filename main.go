// inspired from https://gist.github.com/thedevsaddam/a863148ff9de4cd3fcc628191005ab2e#file-rest_api-go

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/controllers"
	"main/models"
)

func main() {
	// replace by init() inside config package
	config.LoadConfig()

	// set default middleware for logging
	router := gin.Default()

	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
		// allow all origins and methods
		router.Use(cors.Default())
		// todo graceful shutdown
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// replace by init() inside models package
	// connect to db and does the migrations
	models.ConnectDataBase()

	// default routes of APIs
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", controllers.CreateTodo)
		v1.GET("/", controllers.FetchAllTodo)
		v1.GET("/:id", controllers.FetchSingleTodo)
		v1.PATCH("/:id", controllers.UpdateTodo)
		v1.DELETE("/:id", controllers.DeleteTodo)
	}

	// kick start the server and handle the errors, like port under 1024
	if err := router.Run(":3000"); err != nil {
		panic(err)
	}

}
