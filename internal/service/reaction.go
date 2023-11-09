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
	repo3 repository.NotificationPost
}

func NewReaction(repo repository.PostReaction, repo2 repository.CommentReaction, repo3 repository.NotificationPost) *ReactionService {
	return &ReactionService{repo1: repo, repo2: repo2, repo3: repo3}
}

func (s *ReactionService) ReactPost(post_id, user_id, author_id, value int64, username string) error {
	if value != 1 && value != -1 {
		return errors.New("BadRequest")
	}

	fmt.Println("HERE IN service react post")
	postReaction, err := s.repo1.FindReation(post_id, user_id, value)
	if err == sql.ErrNoRows {
		fmt.Println("I am here")
		err = s.repo1.CreateReaction(post_id, user_id, value)
		if author_id != user_id {
			fmt.Println("Before create notify reaction")
			err = s.repo3.CreateNotifyReaction(post_id, user_id, author_id, value, username)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
	} else if err != nil {
		return err
	} else {
		if postReaction.Value == value {
			err = s.repo1.DeleteReaction(post_id, user_id)
			if author_id != user_id {
				err = s.repo3.DeletenNotifyReaction(post_id, user_id, author_id, value)
			}
		} else {
			err = s.repo1.LikePost(post_id, user_id, value)
			if author_id != user_id {
				err = s.repo3.DeletenNotifyReaction(post_id, user_id, author_id, value)
				err = s.repo3.CreateNotifyReaction(post_id, user_id, author_id, value, username)
			}
		}
	}

	return nil
}

func (s *ReactionService) ReactComment(comment_id, user_id, value int64) error {
	if value != 1 && value != -1 {
		return errors.New("BadRequest")
	}
	postReaction, err := s.repo2.FindReation(comment_id, user_id, value)
	if err == sql.ErrNoRows {
		err = s.repo2.CreateReaction(comment_id, user_id, value)
	} else if err != nil {
		return err
	} else {
		if postReaction.Value == value {
			err = s.repo2.DeleteReaction(comment_id, user_id)
		} else {
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

func (s *ReactionService) GetCommentReaction(user_id, commentId int64) (int64, error) {
	return s.repo2.GetCommentReaction(user_id, commentId)
}
