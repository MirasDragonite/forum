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

func (r *Auth) CreateUser(user *structs.User) (int64, error) {
	query := `INSERT INTO users(username,email,hash_password) VALUES($1,$2,$3) RETURNING id`

	result, err := r.db.Exec(query, user.GetUserName(), user.GetUserEmail(), user.GetUserHashPassword())
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(email string) (int64, string, error) {
	var id int64
	var hash_password string

	query := "SELECT id, hash_password FROM users WHERE email=$1 "
	row := r.db.QueryRow(query, email)
	err := row.Scan(&id, &hash_password)
	if err != nil {
		return 0, "", err
	}
	return id, hash_password, nil
}

