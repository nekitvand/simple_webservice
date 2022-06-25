package service

import (
	"webservice/internals/repository"
)

type ToDoService struct {
	repository *repository.ToDoRepository
}

func NewUsersProcessor(storage *repository.ToDoRepository) *ToDoService {
	processor := new(ToDoService)
	processor.repository = storage
	return processor
}