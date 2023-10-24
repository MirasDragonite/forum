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
	// if len(topics) == 0 {
	// 	return nil, nil
	// }

	topicSet := make(map[string]struct{})
	for _, ch := range topics {
		topicSet[ch] = struct{}{}
	}

	posts, err := f.repo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	var filtered []structs.Post

	for _, post := range posts {
		if len(topics) == 0 {
			fmt.Println("Here in case")
			if len(post.Topic) == 1 && strings.TrimSpace(post.Topic[0]) == "" {
				fmt.Println("Find post")
				filtered = append(filtered, post)
				continue
			}
		}
		fmt.Println("POST", post.Topic, "Len", len(post.Topic))
		for _, topic := range post.Topic {
			if strings.TrimSpace(topic) == "" {
				continue
			}
			if _, exists := topicSet[topic]; exists {
				filtered = append(filtered, post)
				break
			}
		}
	}

	return filtered, nil
}
