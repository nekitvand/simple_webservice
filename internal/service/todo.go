package service

import (
	"errors"
	"webservice/internal/model"
	"webservice/internal/repository"
)

type ToDoService struct {
	repository *repository.ToDoRepository
}

func NewUsersService(repository *repository.ToDoRepository) *ToDoService {
	processor := new(ToDoService)
	processor.repository = repository
	return processor
}

func (service *ToDoService) CreateToDo (todo model.ToDoModel) error {
	if todo.Title == "" {
		return errors.New("name не может быть пустым")
	}
	if todo.Text == "" {
		return errors.New("text не может быть пустым")

	}
	return service.repository.CreateToDo(todo)
}