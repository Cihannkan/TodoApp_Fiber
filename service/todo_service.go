package service

import (
	"TodoApp/model"
	"TodoApp/repository"
)

type TodoService struct {
	Repo *repository.TodoRepository
}

func (s *TodoService) CreateTodo(todo *model.Todo) error {
	return s.Repo.Create(todo)
}
