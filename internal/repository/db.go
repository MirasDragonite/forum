package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}

	query := `DROP TABLE IF EXISTS users;
	CREATE TABLE users(id INTEGER PRIMARY KEY, username TEXT, email TEXT,hash_password TEXT);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfuly connect to database")
	return db, nil
}
