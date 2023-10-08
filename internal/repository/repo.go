package repository

import (
	"database/sql"
	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) error
	GetUserBy(element, from string) (structs.User, error)
	GetSession(userID int64) (structs.Session, error)
	CreateToken(user structs.User, token, expaired_data string) error
	UpdateToken(user structs.User, token, expaired_data string) error
	DeleteToken(token string) error
}

type PostRedact interface {
	CreatePost(post *structs.Post) error
	GetUSerID(token string) (int64, error)
	// DislikePost()
	// RedactContentPost()
	// DeletePost()
}

type Repository struct {
	Authorization
	PostRedact
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: NewAuth(db), PostRedact: NewPostRedactDB(db)}
}
