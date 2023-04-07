package mock_repo

import (
	"todo-backend/entity"

	"github.com/stretchr/testify/mock"
)

type MockTodoRepo struct {
	mock.Mock
}

func (ts *MockTodoRepo) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	args := ts.Called(todo)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (ts *MockTodoRepo) GetAllTodos() ([]*entity.Todo, error) {
	args := ts.Called()
	return args.Get(0).([]*entity.Todo), args.Error(1)
}

func (ts *MockTodoRepo) GetTodoById(id string) (*entity.Todo, error) {
	args := ts.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (ts *MockTodoRepo) UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error) {
	args := ts.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (ts *MockTodoRepo) DeleteTodoById(id string) (*entity.Todo, error) {
	args := ts.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Todo), args.Error(1)
}
