package api

import (
	"webservice/internal/handler"
	"github.com/gorilla/mux"
)


func CreateRoutes(toDoHandler *handler.ToDoHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todo/create", toDoHandler.CreateToDo).Methods("POST")
	return r
}
