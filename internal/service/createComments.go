package service

import (
	"fmt"
	"forum/internal/repository"
	"forum/structs"
	"strings"
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
		return nil
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
		fmt.Println("Ch.Commentid:", ch.CommentID)
		likes, dislikes, err = repo.reaction.AllReactions(ch.CommentID)

		if err != nil {
			fmt.Println(err.Error)
			return nil, err
		}

		comments[i].Like = likes
		comments[i].Dislike = dislikes
		// ch.Like = likes
		// ch.Dislike = dislikes

	}
	return comments, nil
}
