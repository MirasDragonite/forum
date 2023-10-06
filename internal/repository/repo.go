package repository

import (
	"database/sql"

	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) (int64, error)
	GetUser(email string) (int64, string, error)
}

type Repository struct {
	Authorization
}

type PostRedact interface {
	CreatePost()
	LikePost(LikePostinadoiad())
	DislikePost()
	WriteCommentPost() {
		db(query)
	}
	RedactContentPost()
	DeletePost()
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: NewAuth(db)}
}
