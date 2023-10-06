package service

import (
	"errors"

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

func (s *Auth) CreateUser(user *structs.User) (int64, error) {
	hashPassword, err := hashPassword(user.GetUserHashPassword())
	if err != nil {
		return 0, err
	}

	user.ChangeUserHashPassword(hashPassword)

	return s.repo.CreateUser(user)
}

func (s *Auth) GetUser(email, password string) (int64, error) {
	id, hash_password, err := s.repo.GetUser(email)

	if checkPasswordHash(password, hash_password) {
		return id, err
	} else {
		return 0, errors.New("Passwords not compatible")
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
