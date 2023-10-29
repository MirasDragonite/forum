package service

import "forum/internal/repository"

type Notify struct {
	repoPostRect repository.PostReaction
	repoPost     repository.PostRedact
}

func NewNotify(repoPostRect repository.PostReaction, repoPost repository.PostRedact) *Notify {
	return &Notify{repoPost: repoPost, repoPostRect: repoPostRect}
}

func (s *Notify) AllUserNotifications(userID int64) error {
	// userPosts, err := s.repoPost.GetAllUserPosts(userID)
	// if err != nil {
	// 	return err
	// }

	return nil
}
