package service

import "forum/internal/repository"

type Authorization interface{}

type Service struct {
	Authorization
}

func NewService(*repository.Repository) *Service {
	return &Service{}
}
