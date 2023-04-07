package repo

import (
	"errors"
	"time"
	"todo-backend/entity"
	"todo-backend/service"

	"github.com/google/uuid"
)

type TodoRepositoryImpl struct {
	todos []*entity.Todo
}

func CreateTodoRepo() service.TodoRepo {
	return &TodoRepositoryImpl{make([]*entity.Todo, 0)}
}

func (repo *TodoRepositoryImpl) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	if todo.Content == "" {
		return nil, errors.New("Content is empty")
	}

	id := uuid.New()
	todo.ID = id.String()
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
		if todo.ID == id {
			return todo, nil
		}
	}

	return nil, errors.New("Todo not found")
}

func (repo *TodoRepositoryImpl) UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error) {
	for _, todoItem := range repo.todos {
		if todoItem.ID == id {
			todoItem.Content = todo.Content
			todoItem.UpdateAt = time.Now()
			return todoItem, nil
		}
	}

	return nil, errors.New("Todo not found")
}

func (repo *TodoRepositoryImpl) DeleteTodoById(id string) (*entity.Todo, error) {
	for i, todoItem := range repo.todos {
		if todoItem.ID == id {
			repo.todos = append(repo.todos[:i], repo.todos[i+1:]...)
			return todoItem, nil
		}
	}

	return nil, errors.New("Todo not found")
}
