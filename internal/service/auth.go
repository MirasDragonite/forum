package service

import (
	"forum/internal/repository"
	"forum/structs"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) CreateUser(user structs.User) (int64, error) {
	hashPassword, err := HashPassword(user.HashedPassword)
	if err != nil {
		return 0, err
	}
	user.HashedPassword = hashPassword

	return s.repo.CreateUser(user)
}

func (s *Auth) GetUser(name, password string) (int64, error) {
	return s.repo.GetUser(name, password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
