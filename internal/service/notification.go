package service

import (
	"forum/internal/repository"
	"forum/structs"
)

type Notify struct {
	repo repository.NotificationPost
}

func NewNotify(repo repository.NotificationPost) *Notify {
	return &Notify{repo: repo}
}

func (s *Notify) AllUserNotifications(author_id int64) ([]structs.Notify, error) {
	return s.repo.GetPostNotification(author_id)
}
