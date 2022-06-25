package service

import (
	"webservice/internal/repository"
)

type ToDoService struct {
	repository *repository.ToDoRepository
}

func NewUsersService(storage *repository.ToDoRepository) *ToDoService {
	processor := new(ToDoService)
	processor.repository = storage
	return processor
}