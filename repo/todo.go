package repo

import (
	"fmt"
	"todo-backend/entity"
	"todo-backend/service"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	DB *gorm.DB
}

func CreateTodoRepo(DB *gorm.DB) service.TodoRepo {
	return &TodoRepositoryImpl{DB}
}

func (repo *TodoRepositoryImpl) CreateTodo(todo *entity.Todo) (*entity.Todo, error) {
	err := repo.DB.Create(&todo).Error
	if err != nil {
		fmt.Printf("[TodoRepoImpl.CreateTodo] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}

	return todo, nil
}

func (repo *TodoRepositoryImpl) GetAllTodos() ([]*entity.Todo, error) {
	var todos []*entity.Todo

	err := repo.DB.Find(&todos).Error
	if err != nil {
		fmt.Printf("[TodoRepoImpl.GetAllTodos] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}

	return todos, nil
}

func (repo *TodoRepositoryImpl) GetTodoById(id string) (*entity.Todo, error) {
	var todo = entity.Todo{}

	err := repo.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		fmt.Printf("[TodoRepoImpl.GetTodoById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}

	return &todo, nil
}

func (repo *TodoRepositoryImpl) UpdateTodoById(id string, todo *entity.Todo) (*entity.Todo, error) {
	var upTodo = entity.Todo{}

	err := repo.DB.Where("id = ?", id).First(&upTodo, id).Updates(todo).Error
	if err != nil {
		fmt.Printf("[TodoRepoImpl.UpdateTodoById] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}

	return &upTodo, nil
}

func (repo *TodoRepositoryImpl) DeleteTodoById(id string) error {
	var todo = entity.Todo{}

	err := repo.DB.Where("id = ?", id).First(&todo).Delete(&todo).Error
	if err != nil {
		fmt.Printf("[TodoRepoImpl.DeleteTodoById] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}

	return nil
}
