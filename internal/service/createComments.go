package service

import (
	"forum/internal/repository"
	"forum/structs"
)

type CommentRed struct {
	repo repository.CommentRedact
}

func NewCommentRed(repo repository.CommentRedact) *CommentRed {
	return &CommentRed{repo: repo}
}

func (repo *CommentRed) CreateComment(comm *structs.Comment) error {
	return repo.repo.CreateComment(comm)
}
