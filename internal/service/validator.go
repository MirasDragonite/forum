package service

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"forum/structs"
)

func checkUserInput(user *structs.User) error {
	if err := isEmailValid(user.Email); err != nil {
		return err
	}
	if err := isNameValid(user.Username); err != nil {
		return err
	}
	if err := isPasswordValid(user.HashedPassword); err != nil {
		return nil
	}
	return nil
}

func isEmailValid(email string) error {
	fmt.Println("email", email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Not valid email")
	}
	if strings.TrimSpace(email) == "" {
		return errors.New("User email can't be empty")
	}
	if len(email) < 2 || len(email) > 32 {
		return errors.New("The length of the email is not up to standard ")
	}
	return nil
}

func isNameValid(name string) error {
	if len(strings.TrimSpace(name)) < 2 || len(strings.TrimSpace(name)) > 32 {
		return errors.New("The length of the user is not up to standard ")
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("User name can't be empty")
	}
	return nil
}

func isPasswordValid(password string) error {
	if len(strings.TrimSpace(password)) < 2 || len(strings.TrimSpace(password)) > 32 {
		return errors.New("The length of the password is not up to standard ")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("User password can't be empty")
	}
	return nil
}
