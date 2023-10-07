package service

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"forum/internal/repository"
	"forum/structs"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userNameDataBaseName = "username"
	emailDataBaseName    = "email"
	timeFormat           = "2006-01-02 15:04:05"
)

type Auth struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *Auth {
	return &Auth{repo: repo}
}

func (s *Auth) CreateUser(user *structs.User) (int64, error) {
	hashPassword, err := hashPassword(user.HashedPassword)
	if err != nil {
		return 0, err
	}

	user.HashedPassword = hashPassword

	return s.repo.CreateUser(user)
}

func (s *Auth) GetUser(email, password string) (*http.Cookie, error) {
	user, err := s.repo.GetUserBy(email, emailDataBaseName)

	if checkPasswordHash(password, user.HashedPassword) {
	} else {
		return nil, errors.New("Passwords not compatible")
	}

	session, err := s.repo.GetSession(user.Id)

	cookie := http.Cookie{Name: "Token"}
	expiration := giveExpirationData()
	hashedToken := createToken()
	expirationInStringFormat := expiration.Format(timeFormat)

	if err == sql.ErrNoRows {
		session.ExpairedData = expirationInStringFormat
		s.repo.CreateToken(user, hashedToken, session.ExpairedData)
		cookie.Value = hashedToken
		cookie.Expires = expiration
		return &cookie, nil

	} else {
		if session.ExpairedData < time.Now().Format(timeFormat) {
			s.repo.UpdateToken(user, hashedToken, expirationInStringFormat)
			return &cookie, nil
		}

		parsedTime, err := time.Parse(timeFormat, session.ExpairedData)
		if err != nil {
			return nil, err
		}
		cookie.Value = session.Token
		cookie.Expires = parsedTime

	}

	return &cookie, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken() string {
	token := uuid.NewV4()
	return token.String()
}

func giveExpirationData() time.Time {
	livingTime := 60 * time.Minute
	expiration := time.Now().Add(livingTime)
	return expiration
}
