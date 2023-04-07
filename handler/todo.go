package handler

import (
	"net/http"
	"todo-backend/entity"

	"github.com/gin-gonic/gin"
)

type TodoRepo interface {
	CreateTodo(todo *entity.Todo) (*entity.Todo, error)
	GetAllTodos() ([]*entity.Todo, error)
	GetTodoById(id string) (*entity.Todo, error)
	UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error)
	DeleteTodoById(id string) error
}

type TodoHandler struct {
	TodoService TodoRepo
}

func CreateTodoHandler(todoService TodoRepo) TodoHandler {
	return TodoHandler{todoService}
}

func (th *TodoHandler) CreateTodo(ctx *gin.Context) {
	var todo entity.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := th.TodoService.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data created": newTodo})
}

func (th *TodoHandler) GetAllTodos(ctx *gin.Context) {
	todos, err := th.TodoService.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"datas": todos})
}

func (th *TodoHandler) GetTodoById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	todo, err := th.TodoService.GetTodoById(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todo})
}

func (th *TodoHandler) UpdateTodoById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	_, err := th.TodoService.GetTodoById(idStr)
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

	todo, err := th.TodoService.UpdateTodoById(idStr, &tempTodo)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data updated": todo})
}

func (th *TodoHandler) DeleteTodoById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	_, err := th.TodoService.GetTodoById(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = th.TodoService.DeleteTodoById(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted successfully item:": idStr})
}
