package repository

import (
	"database/sql"

	"forum/structs"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user structs.User) (int64, error) {
	query := `INSERT INTO users(username,email,hash_password) VALUES($1,$2,$3) RETURNING id`

	result, err := r.db.Exec(query, user.Username, user.Email, user.HashedPassword)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(name, password string) (int64, error) {
	var id int64
	query := "SELECT id FROM users WHERE username=$1 AND password=$2"
	row := r.db.QueryRow(query, name, password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
