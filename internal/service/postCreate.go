package service

import (
	"forum/internal/repository"
	"forum/structs"
)

type PostRed struct {
	repo repository.PostRedact
}

func NewPostRed(repo repository.PostRedact) *PostRed {
	return &PostRed{repo: repo}
}

func (repo *PostRed) CreatePost(post *structs.Post, token string) error {
	userID, err := repo.repo.GetUSerID(token)
	if err != nil {
		return err
	}
	post.PostAuthorID = userID
	post.Like = 0
	post.Dislike = 0
	post.Comments = []structs.Comment{}
	err = repo.repo.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}


