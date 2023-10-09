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
	GetSessionByToken(token string) (structs.Session, error)
	GetUserById(id int64) (structs.User, error)
}

// Post actions
type PostRedact interface {
	CreatePost(post *structs.Post) error
	GetUSerID(token string) (int64, error)
	GetUserName(userID int64) (string, error)
	GetPostBy(from, value string) (*structs.Post, error)
	LikePost(post *structs.Post) error
	DislikePost(post *structs.Post) error
	RedactContentPost(post *structs.Post) error
	DeletePost(post *structs.Post) error
}

type CommentRedact interface {
	CreateComment(comm *structs.Comment) error
}

type Repository struct {
	Authorization
	PostRedact
	CommentRedact
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Authorization: NewAuth(db), PostRedact: NewPostRedactDB(db), CommentRedact: NewCommentRedactDB(db)}
}
