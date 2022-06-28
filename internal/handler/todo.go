package handler

import (
	"encoding/json"
	"net/http"
	"webservice/internal/model"
	"webservice/internal/service"
)


type ToDoHandler struct {
	service *service.ToDoService
}

func NewToDoHandler(service *service.ToDoService) *ToDoHandler {
	handler := new(ToDoHandler)
	handler.service = service
	return handler
}

func (handler *ToDoHandler) CreateToDo(w http.ResponseWriter, r *http.Request) {
	var newToDo model.ToDoModel

	err := json.NewDecoder(r.Body).Decode(&newToDo)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.service.CreateToDo(newToDo)
	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{} {
		"result": "OK",
		"info" : "created",
	}

	WrapOK(w,response)
}	


func (handler *ToDoHandler) GetAllToDo(w http.ResponseWriter, r *http.Request) {

	todo_list, err := handler.service.GetAllToDo()
	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{} {
		"result": todo_list,
	}

	WrapOK(w,response)
}	


