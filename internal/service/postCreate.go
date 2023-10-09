package service

import (
	"errors"
	"fmt"
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
	fmt.Println("Service:" , post)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRed) GetUSerID(token string) (int64, error) {
	return repo.repo.GetUSerID(token)
}

func (repo *PostRed) GetPostBy(from, value string) (*structs.Post, error) {
	return repo.repo.GetPostBy(from, value)
}

func (repo *PostRed) GetUserName(token string) (string, error) {
	userID, err := repo.repo.GetUSerID(token)
	if err != nil {
		return "", errors.New("No such user with this token...")
	}

	userName, err := repo.repo.GetUserName(userID)
	return userName, nil

}

func (repo *PostRed) LikePost(post *structs.Post) error {
	if post.Like == 0 && post.Dislike == 0 {
		post.Like++
	} else if post.Like == 1 && post.Dislike == 0 {
		post.Like = 1
	} else if post.Like == 0 && post.Dislike == 1 {
		post.Like++
		post.Dislike--
	} else {
		return errors.New("Error in service/LikePost(). Like == 1 && Dislike == 1")
	}
	return repo.repo.LikePost(post)
}

func (repo *PostRed) DislikePost(post *structs.Post) error {
	if post.Like == 0 && post.Dislike == 0 {
		post.Dislike++
	} else if post.Like == 1 && post.Dislike == 0 {
		post.Like--
		post.Dislike++
	} else if post.Like == 0 && post.Dislike == 1 {
		post.Dislike = 1
	} else {
		return errors.New("Error in service/LikePost(). Like == 1 && Dislike == 1")
	}
	return repo.repo.DislikePost(post)
}

func (repo *PostRed) RedactContentPost(post *structs.Post, newContent string) error {
	post.Content = newContent
	return repo.repo.RedactContentPost(post)
}

func (repo *PostRed) DeletePost(post *structs.Post) error {
	return repo.repo.DeletePost(post)
}
