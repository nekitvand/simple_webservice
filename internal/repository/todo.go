package repository

import (
	"github.com/jmoiron/sqlx"
)

type ToDoRepository struct {
	db *sqlx.DB
}


func NewToDoRepository(db *sqlx.DB) *ToDoRepository {
	repository := new(ToDoRepository)
	repository.db = db
	return repository

}