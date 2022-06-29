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
	service := new(ToDoService)
	service.repository = repository
	return service
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

func (service *ToDoService) GetAllToDo() ([]model.ToDoModel,error) {
	todo_list,err := service.repository.GetAllToDo()
	
	if err != nil {
		return nil, errors.New("ошибка получения данных")
	}
	if len(todo_list) == 0 {
		return todo_list, errors.New("нет записей в базе")
	}

	return todo_list, nil
}

func (service *ToDoService) UpdateFieldToDo(todo model.ToDoModel) error {
	return service.repository.UpdateFieldToDo(todo)
}