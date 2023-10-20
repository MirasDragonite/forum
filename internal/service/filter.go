package service

import (
	"fmt"
	"strings"

	"forum/internal/repository"
	"forum/structs"
)

type Filtering struct {
	repo repository.PostRedact
}

func NewFilter(repo repository.PostRedact) *Filtering {
	return &Filtering{repo: repo}
}

func (f *Filtering) Filter(topics []string) ([]structs.Post, error) {
	// tops := make(map[string]bool)
	// for _, ch := range topics {
	// 	tops[ch] = true
	// }
	fmt.Println()
	if len(topics) == 0 {
		return nil, nil
	}

	posts, err := f.repo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	var filtered []structs.Post
	for _, post := range posts {
		for _, topic := range post.Topic {
			if strings.TrimSpace(topic) == "" {
				continue
			}
			for _, ch := range topics {
				if topic == ch {
					fmt.Println(topics)
					fmt.Println(post.Topic)
					fmt.Println("Topic", topic)
					fmt.Println("CH:", ch)
					filtered = append(filtered, post)
					break
				}
			}
		}
	}

	return filtered, nil
}
