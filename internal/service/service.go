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
}

type PostRedact interface {
	CreatePost(post *structs.Post, token string) error
	// DislikePost()
	// WriteCommentPost()
	// RedactContentPost()
	// DeletePost()
}

type Service struct {
	Authorization
	PostRedact
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuth(repo.Authorization), PostRedact: NewPostRed(repo.PostRedact)}
}
