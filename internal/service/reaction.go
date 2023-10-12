package service

import (
	"database/sql"
	"errors"
	"fmt"

	"forum/internal/repository"
)

type ReactionService struct {
	repo repository.Reaction
}

func NewReaction(repo repository.Reaction) *ReactionService {
	return &ReactionService{repo: repo}
}

func (s *ReactionService) ReactPost(post_id, user_id, value int64) error {
	if value != 1 && value != -1 {
		return errors.New("BadRequest")
	}
	postReaction, err := s.repo.FindReation(post_id, user_id, value)
	if err == sql.ErrNoRows {
		fmt.Println("Creating")
		err = s.repo.CreateReaction(post_id, user_id, value)
	} else if err != nil {
		return err
	} else {
		if postReaction.Value == value {
			fmt.Println("Delete")
			err = s.repo.DeleteReaction(post_id, user_id)
		} else {
			fmt.Println("React")
			err = s.repo.LikePost(post_id, user_id, value)
		}
	}

	return nil
}

func (s *ReactionService) AllReactions(post_id int64) (int64, int64, error) {
	return s.repo.AllReactions(post_id)
}
