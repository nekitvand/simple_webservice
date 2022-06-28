package repository

import (
	"fmt"
	"webservice/internal/model"

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

func (repository *ToDoRepository) CreateToDo (todo model.ToDoModel) error {
	_, err := repository.db.Exec("INSERT INTO todo(title,text) VALUES ($1,$2)",todo.Title,todo.Text)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (repository *ToDoRepository) GetAllToDo() ([]model.ToDoModel,error) {
	query := "SELECT id,title,text FROM todo"

	var result []model.ToDoModel

	err := repository.db.Select(&result, query)
	

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil

}