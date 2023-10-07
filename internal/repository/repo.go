package repository

import (
	"database/sql"

	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) (int64, error)
	GetUserBy(element, from string) (structs.User, error)
	GetSession(userID int64) (structs.Session, error)
	CreateToken(user structs.User, token, expaired_data string) error
	UpdateToken(user structs.User, token, expaired_data string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: NewAuth(db)}
}
