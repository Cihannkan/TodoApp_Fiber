package repository

import (
	"TodoApp/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.DB.Create(todo).Error
}
