package service

import (
	"todo-backend/entity"
)

type TodoRepo interface {
	CreateTodo(todo *entity.Todo) (*entity.Todo, error)
	GetAllTodos() ([]*entity.Todo, error)
	GetTodoById(id string) (*entity.Todo, error)
	UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error)
	DeleteTodoById(id string) error
}

type TodoService struct {
	TodoRepo TodoRepo
}

func CreateTodoService(todoRepo TodoRepo) TodoRepo {
	return &TodoService{todoRepo}
}

func (ts *TodoService) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	return ts.TodoRepo.CreateTodo(todo)
}

func (ts *TodoService) GetAllTodos() ([]*entity.Todo, error) {
	return ts.TodoRepo.GetAllTodos()
}

func (ts *TodoService) GetTodoById(id string) (*entity.Todo, error) {
	return ts.TodoRepo.GetTodoById(id)
}

func (ts *TodoService) UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error) {
	return ts.TodoRepo.UpdateTodoById(id, todo)
}

func (ts *TodoService) DeleteTodoById(id string) error {
	return ts.TodoRepo.DeleteTodoById(id)
}
