package service

import (
	"forum/internal/repository"
	"forum/structs"
)

type Filtering struct {
	repo repository.PostRedact
}

func NewFilter(repo repository.PostRedact) *Filtering {
	return &Filtering{repo: repo}
}

func (f *Filtering) Filter(java, kotlin, python, topic string) ([]structs.Post, error) {
	return f.repo.GetFilteredPosts(java, kotlin, python, topic)
}
