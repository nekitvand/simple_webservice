package api

import (
	"webservice/internal/handler"
	"github.com/gorilla/mux"
)


func CreateRoutes(toDoHandler *handler.ToDoHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todo/create", toDoHandler.CreateToDo).Methods("POST")
	r.HandleFunc("/todo/get_all", toDoHandler.GetAllToDo).Methods("GET")
	r.HandleFunc("/todo/update/{id:[0-9]+}", toDoHandler.UpdateToDo).Methods("PUT")
	r.HandleFunc("/todo/update_field/{id:[0-9]+}", toDoHandler.UpdateFieldToDo).Methods("PATCH")
	r.HandleFunc("/todo/delete/{id:[0-9]+}", toDoHandler.DeleteToDo).Methods("DELETE")
	return r
}