package service

import (
	"errors"
	"strings"

	"forum/internal/repository"
	"forum/structs"
)

type CommentRed struct {
	repo     repository.CommentRedact
	reaction repository.CommentReaction
}

func NewCommentRed(repo repository.CommentRedact, reaction repository.CommentReaction) *CommentRed {
	return &CommentRed{repo: repo, reaction: reaction}
}

func (repo *CommentRed) CreateComment(comm *structs.Comment) error {
	if strings.TrimSpace(comm.Content) == "" {
		return errors.New("Can't be empty")
	}
	return repo.repo.CreateComment(comm)
}

func (repo *CommentRed) GetCommentByID(commentID int64) (structs.Comment, error) {
	return repo.repo.GetCommentByID(commentID)
}

func (repo *CommentRed) GetAllComments(postID, userID int64) ([]structs.Comment, error) {
	comments, err := repo.repo.GetAllComments(postID, userID)
	if err != nil {
		return nil, err
	}

	for i, ch := range comments {

		var likes int64
		var dislikes int64

		likes, dislikes, err = repo.reaction.AllReactions(ch.CommentID)

		if err != nil {
			return nil, err
		}

		comments[i].Like = likes
		comments[i].Dislike = dislikes
		// ch.Like = likes
		// ch.Dislike = dislikes

	}
	return comments, nil
}
