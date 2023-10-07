package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"
	"net/url"
	"time"

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

func (s *Auth) GetUser(email, password string) (*http.Cookie, int64, string, error) {
	id, hash_password, err := s.repo.GetUser(email)

	if checkPasswordHash(password, hash_password) {

		time64 := time.Now().Unix()
		timeInt := string(rune(time64))
		token := email + password + timeInt
		hashToken := md5.Sum([]byte(token))
		hashedToken := hex.EncodeToString(hashToken[:])
		// h.Cache[hashedToken] = id
		livingTime := 60 * time.Minute
		expiration := time.Now().Add(livingTime)
		cookie := http.Cookie{Name: "Token", Value: url.QueryEscape(hashedToken), Expires: expiration}
		return &cookie, id, hashedToken, err
	} else {
		return nil, 0, "", errors.New("Passwords not compatible")
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
