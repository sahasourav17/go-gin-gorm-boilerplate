package main

import (
	"example/explore-go/config"
	"example/explore-go/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.POST("/todos", controllers.AddTodo)
	router.PATCH("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
	router.Run()
}
