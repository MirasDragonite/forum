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
	repoNot  repository.NotificationPost
}

func NewCommentRed(repo repository.CommentRedact, reaction repository.CommentReaction, repoNot repository.NotificationPost) *CommentRed {
	return &CommentRed{repo: repo, reaction: reaction, repoNot: repoNot}
}

func (repo *CommentRed) CreateComment(comm *structs.Comment, user_id int64) error {
	if strings.TrimSpace(comm.Content) == "" {
		return errors.New("Can't be empty")
	}
	err := repo.repo.CreateComment(comm)
	if err != nil {
		return err
	}
	if comm.CommentAuthorID != user_id {
		err = repo.repoNot.CreateNotifyReaction(comm.PostID, comm.CommentAuthorID, user_id, 3, comm.CommentAuthorName)
		if err != nil {
			return err
		}
	}
	return nil
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
