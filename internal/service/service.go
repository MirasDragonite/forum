package service

import (
	"net/http"

	"forum/internal/repository"
	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) error
	GetUser(email, password string) (*http.Cookie, error)
	DeleteToken(cookie *http.Cookie) error
	GetUserByToken(token string) (*structs.User, error)
}

type PostRedact interface {
	CreatePost(post *structs.Post, token string) error
	GetUSerID(token string) (int64, error)
	GetUserName(token string) (string, error)
	GetPostBy(from, value string) (*structs.Post, error)

	LikePost(post *structs.Post) error

	DislikePost(post *structs.Post) error
	// WriteCommentPost()
	RedactContentPost(post *structs.Post, newContent string) error
	DeletePost(post *structs.Post) error
}

type CommentRedact interface {
	CreateComment(comm *structs.Comment) error
}

type Reaction interface {
	ReactPost(post_id, user_id, value int64) error
	AllReactions(post_id int64) (int64, int64, error)
}

type Service struct {
	Authorization
	PostRedact
	CommentRedact
	Reaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuth(repo.Authorization),
		PostRedact:    NewPostRed(repo.PostRedact),
		CommentRedact: NewCommentRed(repo.CommentRedact),
		Reaction:      NewReaction(repo.Reaction),
	}
}
