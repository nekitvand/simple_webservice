package handlers

import (
	"webservice/internals/service"
	"net/http"
)


type ToDoHandler struct {
	service *service.ToDoService
}

func NewToDoHandler(service *service.ToDoService) *ToDoHandler {
	handler := new(ToDoHandler)
	handler.service = service
	return handler
}

func (handler *ToDoHandler) Create(w http.ResponseWriter, r *http.Request) {
	
}


