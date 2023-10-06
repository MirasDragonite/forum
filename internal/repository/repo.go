package repository

import (
	"database/sql"

	"forum/structs"
)



type Authorization interface {
	CreateUser(user *structs.User) (int64, error)
	GetUser(email string) (int64, string, error)
}



type PostRedact interface {
	CreatePost()
	// LikePost(LikePostinadoiad())
	DislikePost()
	RedactContentPost()
	DeletePost()
}

type Repository struct {
	Authorization
	PostRedact
	
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: NewAuth(db)}
}
