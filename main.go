package main

import (
	"todo-backend/handler"
	"todo-backend/repo"
	"todo-backend/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoRepo := repo.CreateTodoRepo()
	todoService := service.CreateTodoService(todoRepo)
	todoHandler := handler.CreateTodoHandler(router, todoService)

	{
		router.POST("/items", todoHandler.CreateTodo)
		router.GET("/items", todoHandler.GetAllTodos)
		router.GET("/items/:id", todoHandler.GetTodoById)
		router.PUT("/items/:id", todoHandler.UpdateTodoById)
		router.DELETE("/items/:id", todoHandler.DeleteTodoById)
	}

	router.Run()
}
