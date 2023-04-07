package main

import (
	"todo-backend/db"
	"todo-backend/handler"
	"todo-backend/repo"
	"todo-backend/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	DB := db.Init()
	router := gin.Default()

	todoRepo := repo.CreateTodoRepo(DB)
	todoService := service.CreateTodoService(todoRepo)
	todoHandler := handler.CreateTodoHandler(todoService)

	{
		router.POST("/items", todoHandler.CreateTodo)
		router.GET("/items", todoHandler.GetAllTodos)
		router.GET("/items/:id", todoHandler.GetTodoById)
		router.PUT("/items/:id", todoHandler.UpdateTodoById)
		router.DELETE("/items/:id", todoHandler.DeleteTodoById)
	}

	router.Run()
}
