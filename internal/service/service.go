package service

import (
	"forum/internal/repository"
	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) (int64, error)
	GetUser(email, password string) (int64, error)
}
type PostRedact interface {
	CreatePost()
	DislikePost()
	WriteCommentPost()
	RedactContentPost()
	DeletePost()
}

type Service struct {
	Authorization
	PostRedact
}

func NewServiceAuth(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuth(repo.Authorization)}
}

func NewServicePostRedeact(repo *repository.PostRedact) *Service {
	return &Service{PostRedact: NewAuth(repo.PostRedact)}
}