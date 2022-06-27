package migrations

import (
	"database/sql"
	goose "github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS todo(
		id serial PRIMARY KEY,
		title VARCHAR (50) UNIQUE NOT NULL,
		text TEXT NOT NULL 
	);`)
	if err != nil {
		return err
	}
	return nil
}

func downInit(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXIST todo;`)
	if err != nil {
		return err
	}
	return nil
}
