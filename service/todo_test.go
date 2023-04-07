package service

import (
	"errors"
	"testing"
	"time"
	"todo-backend/entity"

	mock_repositories "todo-backend/mock/repo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createMockTodoItem(content string) entity.Todo {
	var todo entity.Todo

	todo.ID = 1
	todo.Content = content
	todo.Status = "new"
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	return todo
}

func createMockTodosList() []*entity.Todo {
	mockTodo := createMockTodoItem("mock content")
	mockTodos := []*entity.Todo{&mockTodo}

	return mockTodos
}

func TestTodo_CreateTodo(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodo := createMockTodoItem("mock content")

		todoRepo.On("CreateTodo", mock.Anything).Once().Return(&mockTodo, nil)

		todoService := &TodoService{todoRepo}
		todo, err := todoService.CreateTodo(&mockTodo)

		assert.NotNil(t, todo)
		assert.Nil(t, err)
		assert.Equal(t, todo, &mockTodo)
	})

	t.Run("empty content", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodo := createMockTodoItem("")

		todoRepo.On("CreateTodo", mock.Anything).Once().Return(nil, errors.New("Content is empty"))

		todoService := &TodoService{todoRepo}
		todo, err := todoService.CreateTodo(&mockTodo)

		assert.NotNil(t, err)
		assert.Nil(t, todo)
		assert.Equal(t, errors.New("Content is empty"), err)
	})
}

func TestTodo_GetAllTodos(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodos := createMockTodosList()

		todoRepo.On("GetAllTodos", mock.Anything).Once().Return(mockTodos, nil)

		todoService := &TodoService{todoRepo}
		todos, err := todoService.GetAllTodos()

		assert.Nil(t, err)
		assert.NotNil(t, todos)
		assert.Equal(t, todos, mockTodos)
	})
}

func TestTodo_GetTodoById(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodo := createMockTodoItem("mock content")

		todoRepo.On("GetTodoById", mock.Anything).Once().Return(&mockTodo, nil)

		todoService := &TodoService{todoRepo}
		todo, err := todoService.GetTodoById("mock_id")

		assert.Nil(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, todo, &mockTodo)
	})

	t.Run("id not found", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)

		todoRepo.On("GetTodoById", mock.Anything).Once().Return(nil, errors.New("Todo not found"))

		todoService := &TodoService{todoRepo}
		todo, err := todoService.GetTodoById("mock_id")

		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Todo not found"), err)
	})
}

func TestTodo_UpdateTodoById(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodo := createMockTodoItem("mock content")
		mockUpdatedTodo := createMockTodoItem("mock updated content")

		todoRepo.On("UpdateTodoById", mock.Anything).Once().Return(&mockUpdatedTodo, nil)

		todoService := &TodoService{todoRepo}
		todo, err := todoService.UpdateTodoById("mock_id", &mockTodo)

		assert.Nil(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, todo, &mockUpdatedTodo)
	})

	t.Run("id not found", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodo := createMockTodoItem("mock content")

		todoRepo.On("UpdateTodoById", mock.Anything).Once().Return(nil, errors.New("Todo not found"))

		todoService := &TodoService{todoRepo}
		todo, err := todoService.UpdateTodoById("mock_id", &mockTodo)

		assert.Nil(t, todo)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Todo not found"), err)
	})
}

func TestTodo_DeleteTodoById(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)
		mockTodo := createMockTodoItem("mock content")

		todoRepo.On("DeleteTodoById", mock.Anything).Once().Return(&mockTodo, nil)

		todoService := &TodoService{todoRepo}
		err := todoService.DeleteTodoById("mock_id")

		assert.Nil(t, err)
	})

	t.Run("id not found", func(t *testing.T) {
		todoRepo := new(mock_repositories.MockTodoRepo)

		todoRepo.On("DeleteTodoById", mock.Anything).Once().Return(nil, errors.New("Todo not found"))

		todoService := &TodoService{todoRepo}
		err := todoService.DeleteTodoById("mock_id")

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("Todo not found"), err)
	})
}
