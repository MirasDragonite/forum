package service

import (
	"fmt"

	"forum/internal/repository"
	"forum/structs"
)

type Auth struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) CreateUser(user structs.User) (int64, error) {
	fmt.Println("User in service:", user)
	return s.repo.CreateUser(user)
}

func (s *Auth) GetUser(name, password string) (int64, error) {
	return s.repo.GetUser(name, password)
}
