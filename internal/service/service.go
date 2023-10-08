package service

import (
	"forum/internal/repository"
	"forum/structs"
	"net/http"
)

type Authorization interface {
	CreateUser(user *structs.User) error
	GetUser(email, password string) (*http.Cookie, error)
	DeleteToken(cookie *http.Cookie) error
}

type PostRedact interface {
	CreatePost(post *structs.Post, token string) error
	GetPostBy(from, value string) (*structs.Post, error)
	LikePost(post *structs.Post) error
	DislikePost(post *structs.Post) error
	// WriteCommentPost()
	RedactContentPost(post *structs.Post) error
	DeletePost(post *structs.Post) error
}

type Service struct {
	Authorization
	PostRedact
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuth(repo.Authorization), PostRedact: NewPostRed(repo.PostRedact)}
}
