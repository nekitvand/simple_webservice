package repository

import (
	"errors"
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

func (repository *ToDoRepository) UpdateFieldToDo(todo model.ToDoModel) error {
	var query string
	if todo.Text != "" {
		query += fmt.Sprintf("text = '%v',",todo.Text)
	}
	if todo.Title != "" {
		query +=  fmt.Sprintf("title = '%v'",todo.Title)
	}
	if query == "" {
		return errors.New("нет значений для измнений")
	}
	
	upd := fmt.Sprintf("UPDATE todo SET %s WHERE id = %d",query,todo.Id)

	_, err := repository.db.Exec(upd)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (repository *ToDoRepository) UpdateToDo(todo model.ToDoModel) error {

	_, err := repository.db.Exec("UPDATE todo SET title = $1,text = $2  WHERE id = $3",todo.Title,todo.Text,todo.Id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}


func (repository *ToDoRepository) DeleteToDo(todo model.ToDoModel) error {

	_, err := repository.db.Exec("DELETE FROM todo WHERE id = $1",todo.Id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}