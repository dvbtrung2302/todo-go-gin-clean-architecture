package repo

import (
	"errors"
	"time"
	"todo-backend/entity"

	"github.com/google/uuid"
)

type TodoRepo interface {
	CreateTodo(todo *entity.Todo) (*entity.Todo, error)
	GetAllTodos() ([]*entity.Todo, error)
	GetTodoById(id string) (*entity.Todo, error)
	UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error)
	DeleteTodoById(id string) (*entity.Todo, error)
}

type TodoRepositoryImpl struct {
	todos []*entity.Todo
}

func CreateTodoRepo() TodoRepo {
	return &TodoRepositoryImpl{}
}

func (repo *TodoRepositoryImpl) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	id := uuid.New()
	todo.Id = id.String()
	todo.Status = "new"
	todo.CreatedAt = time.Now()
	todo.UpdateAt = time.Now()

	repo.todos = append(repo.todos, todo)
	return todo, nil
}

func (repo *TodoRepositoryImpl) GetAllTodos() ([]*entity.Todo, error) {
	return repo.todos, nil
}

func (repo *TodoRepositoryImpl) GetTodoById(id string) (*entity.Todo, error) {
	for _, todo := range repo.todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return nil, errors.New("Todo not found")
}

func (repo *TodoRepositoryImpl) UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error) {
	for _, todoItem := range repo.todos {
		if todoItem.Id == id {
			todoItem.Content = todo.Content
			todoItem.UpdateAt = time.Now()
			return todoItem, nil
		}
	}

	return nil, errors.New("Todo not found")
}

func (repo *TodoRepositoryImpl) DeleteTodoById(id string) (*entity.Todo, error) {
	for i, todoItem := range repo.todos {
		if todoItem.Id == id {
			repo.todos = append(repo.todos[:i], repo.todos[i+1:]...)
			return todoItem, nil
		}
	}

	return nil, errors.New("Todo not found")
}
