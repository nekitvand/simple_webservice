package repository

import (
	"github.com/jmoiron/sqlx"
)

type ToDoRepository struct {
	database *sqlx.DB
}
