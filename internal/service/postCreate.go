package service

import (
	"errors"
	"strconv"

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

func (repo *PostRed) GetUSerID(token string) (int64, error) {
	return repo.repo.GetUSerID(token)
}

func (repo *PostRed) GetPostBy(from, value string, userID int64) (*structs.Post, error) {
	return repo.repo.GetPostBy(from, value, userID)
}

func (repo *PostRed) GetUserName(token string) (string, error) {
	userID, err := repo.repo.GetUSerID(token)
	if err != nil {
		return "", errors.New("No such user with this token...")
	}

	userName, err := repo.repo.GetUserName(userID)
	return userName, nil
}

func (repo *PostRed) RedactContentPost(post *structs.Post, newContent string) error {
	post.Content = newContent
	return repo.repo.RedactContentPost(post)
}

func (repo *PostRed) DeletePost(post *structs.Post) error {
	return repo.repo.DeletePost(post)
}

func (repo *PostRed) GetAllPosts() ([]structs.Post, error) {
	return repo.repo.GetAllPosts()
}

func (repo *PostRed) GetAllLikedPosts(user_id int64) ([]structs.Post, error) {
	reactions, err := repo.repo.GetAllLikedPosts(user_id)
	if err != nil {
		return nil, err
	}
	var posts []structs.Post
	for _, ch := range reactions {
		ch_str := strconv.Itoa(int(ch.PostID))
		post, err := repo.repo.GetPostBy("id", ch_str, user_id)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *post)
	}
	return posts, nil
}

func (repo *PostRed) GetAllUserPosts(user_id int64) ([]structs.Post, error) {
	return repo.repo.GetAllUserPosts(user_id)
}
