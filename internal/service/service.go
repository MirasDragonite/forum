package service

import (
	"forum/internal/repository"
	"forum/structs"
)

type Authorization interface {
	CreateUser(user structs.User) (int64, error)
	GetUser(email, password string) (int64, error)
}
type PostRedact interface {
	CreatePost()
	LikePost(LikePostinadoiad())
	DislikePost()
	WriteCommentPost()
	RedactContentPost()
	DeletePost()
}

type Comment interface {
	CreateComment()
	ChangeContent()
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuth(repo.Authorization)}
}
