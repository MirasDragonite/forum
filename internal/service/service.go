package service

import (
	"net/http"

	"forum/internal/repository"
	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) error
	GetUser(email, password string) (*http.Cookie, error)
	DeleteToken(cookie *http.Cookie) error
	GetUserByToken(token string) (*structs.User, error)
}

type PostRedact interface {
	CreatePost(post *structs.Post, token string) error
	GetUSerID(token string) (int64, error)
	GetUserName(token string) (string, error)
	GetPostBy(from, value string, userID int64) (*structs.Post, error)
	GetAllPosts() ([]structs.Post, error)
	RedactContentPost(post *structs.Post, newContent string) error
	DeletePost(post *structs.Post) error
	GetAllLikedPosts(user_id int64) ([]structs.Post, error)
	GetAllUserPosts(user_id int64) ([]structs.Post, error)
	GetAllDislikedPosts(user_id int64) ([]structs.Post, error)
}

type CommentRedact interface {
	CreateComment(comm *structs.Comment, user_id int64) error
	GetAllComments(postID, userID int64) ([]structs.Comment, error)
	GetCommentByID(commentID int64) (structs.Comment, error)
	GetAllUserComments(userID int64) ([]structs.Comment, error)
}

type Reaction interface {
	ReactPost(post_id, user_id, author_id, value int64, username string) error
	AllPostReactions(post_id int64) (int64, int64, error)
	ReactComment(comment_id, user_id, value int64) error
	AllCommentReactions(post_id int64) (int64, int64, error)
	GetPostReaction(user_id, post_id int64) (int64, error)
	GetCommentReaction(user_id, commentId int64) (int64, error)
}
type Filter interface {
	Filter(java, kotlin, python, topic string) ([]structs.Post, error)
}

type Notification interface {
	AllUserNotifications(author_id int64) ([]structs.Notify, error)
}

type Service struct {
	Authorization
	PostRedact
	CommentRedact
	Reaction
	Filter
	Notification
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuth(repo.Authorization),
		PostRedact:    NewPostRed(repo.PostRedact),
		CommentRedact: NewCommentRed(repo.CommentRedact, repo.CommentReaction, repo.NotificationPost),
		Reaction:      NewReaction(repo.PostReaction, repo.CommentReaction, repo.NotificationPost),
		Filter:        NewFilter(repo.PostRedact),
		Notification:  NewNotify(repo.NotificationPost),
	}
}
