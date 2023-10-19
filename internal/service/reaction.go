package service

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/repository"
)

type ReactionService struct {
	repo1 repository.PostReaction
	repo2 repository.CommentReaction
}

func NewReaction(repo repository.PostReaction, repo2 repository.CommentReaction) *ReactionService {
	return &ReactionService{repo1: repo, repo2: repo2}
}

func (s *ReactionService) ReactPost(post_id, user_id, value int64) error {
	if value != 1 && value != -1 {
		return errors.New("BadRequest")
	}
	postReaction, err := s.repo1.FindReation(post_id, user_id, value)
	if err == sql.ErrNoRows {
		fmt.Println("Creating")
		err = s.repo1.CreateReaction(post_id, user_id, value)
	} else if err != nil {
		return err
	} else {
		if postReaction.Value == value {
			fmt.Println("Delete")
			err = s.repo1.DeleteReaction(post_id, user_id)
		} else {
			fmt.Println("React")
			err = s.repo1.LikePost(post_id, user_id, value)
		}
	}

	return nil
}

func (s *ReactionService) ReactComment(comment_id, user_id, value int64) error {
	fmt.Println("REACT COMMENT:", comment_id, user_id, value)
	if value != 1 && value != -1 {
		return errors.New("BadRequest")
	}
	postReaction, err := s.repo2.FindReation(comment_id, user_id, value)
	if err == sql.ErrNoRows {
		fmt.Println("Creating")
		err = s.repo2.CreateReaction(comment_id, user_id, value)
	} else if err != nil {
		return err
	} else {
		if postReaction.Value == value {
			fmt.Println("Delete")
			err = s.repo2.DeleteReaction(comment_id, user_id)
		} else {
			fmt.Println("React")
			err = s.repo2.LikeComment(comment_id, user_id, value)
		}
	}

	return nil
}

func (s *ReactionService) AllPostReactions(post_id int64) (int64, int64, error) {
	return s.repo1.AllReactions(post_id)
}

func (s *ReactionService) AllCommentReactions(post_id int64) (int64, int64, error) {
	return s.repo2.AllReactions(post_id)
}

func (s *ReactionService) GetPostReaction(user_id, post_id int64) (int64, error) {
	return s.repo1.GetPostReaction(user_id, post_id)
}
