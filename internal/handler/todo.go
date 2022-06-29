package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webservice/internal/model"
	"webservice/internal/service"

	"github.com/gorilla/mux"
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


func (handler *ToDoHandler) UpdateFieldToDo(w http.ResponseWriter, r *http.Request) {
	var newToDo model.ToDoModel

	err := json.NewDecoder(r.Body).Decode(&newToDo)
	newToDo.Id,_ = strconv.Atoi(mux.Vars(r)["id"])
	
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.service.UpdateFieldToDo(newToDo)
	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{} {
		"result": "OK",
		"info" : "update todo " + strconv.Itoa(newToDo.Id),
	}

	WrapOK(w,response)
}

func (handler *ToDoHandler) UpdateToDo(w http.ResponseWriter, r *http.Request) {
	var newToDo model.ToDoModel

	err := json.NewDecoder(r.Body).Decode(&newToDo)
	newToDo.Id,_ = strconv.Atoi(mux.Vars(r)["id"])
	
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.service.UpdateToDo(newToDo)
	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{} {
		"result": "OK",
		"info" : "update todo " + strconv.Itoa(newToDo.Id),
	}

	WrapOK(w,response)
}

func (handler *ToDoHandler) DeleteToDo(w http.ResponseWriter, r *http.Request) {
	var newToDo model.ToDoModel

	newToDo.Id,_ = strconv.Atoi(mux.Vars(r)["id"])
	
	err := handler.service.DeleteToDo(newToDo)
	if err != nil {
		WrapError(w, err)
		return
	}

	var response = map[string]interface{} {
		"result": "OK",
		"info" : "delete todo " + strconv.Itoa(newToDo.Id),
	}

	WrapOK(w,response)
}