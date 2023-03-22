package main

import (
	"net/http"
	"todo-backend/entity"
	"todo-backend/repo"
	"todo-backend/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoRepo := repo.CreateTodoRepo()
	todoService := service.CreateTodoService(todoRepo)

	{
		router.POST("/items", createTodo(todoService))
		router.GET("/items", getAllTodos(todoService))
		router.GET("/items/:id", getTodoById(todoService))
		router.PUT("/items/:id", updateTodoById(todoService))
		router.DELETE("/items/:id", deleteTodoById(todoService))
	}

	router.Run()
}

func createTodo(todoService service.TodoRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var todo entity.Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newTodo, err := todoService.CreateTodo(&todo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data created": newTodo})
	}
}

func getAllTodos(todoService service.TodoRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todos, err := todoService.GetAllTodos()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"datas": todos})
	}
}

func getTodoById(todoService service.TodoRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		todo, err := todoService.GetTodoById(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": todo})
	}
}

func updateTodoById(todoService service.TodoRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		_, err := todoService.GetTodoById(idStr)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var tempTodo entity.Todo
		if err := ctx.ShouldBindJSON(&tempTodo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if tempTodo.Content == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "content cannot be empty"})
			return

		}

		todo, err := todoService.UpdateTodoById(idStr, &tempTodo)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data updated": todo})
	}
}

func deleteTodoById(todoService service.TodoRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")

		_, err := todoService.GetTodoById(idStr)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		todo, err := todoService.DeleteTodoById(idStr)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data deleted": todo})
	}
}
