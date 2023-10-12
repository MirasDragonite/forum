package service

import (
	"database/sql"

	"forum/internal/repository"
)

type ReactionService struct {
	repo repository.Reaction
}

func NewReaction(repo repository.Reaction) *ReactionService {
	return &ReactionService{repo: repo}
}

func (s *ReactionService) ReactPost(post_id, user_id, value int64) error {
	postReaction, err := s.repo.FindReation(post_id, user_id, value)
	if err == sql.ErrNoRows {
		err = s.repo.CreateReaction(post_id, user_id, value)
	} else if err != nil {
		return err
	} else {
		if postReaction.Value == value {
			err = s.repo.DeleteReaction(post_id, user_id)
		} else {
			err = s.repo.LikePost(post_id, user_id, value)
		}
	}

	return nil
}
