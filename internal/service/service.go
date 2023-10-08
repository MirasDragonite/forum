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

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuth(repo.Authorization)}
}
