package repository

import (
	"database/sql"
	"forum/structs"
)

type Authorization interface {
	CreateUser(user *structs.User) error
	GetUserBy(element, from string) (structs.User, error)
	GetSession(userID int64) (structs.Session, error)
	CreateToken(user structs.User, token, expaired_data string) error
	UpdateToken(user structs.User, token, expaired_data string) error
	DeleteToken(token string) error
	GetSessionByToken(token string) (structs.Session, error)
	GetUserById(id int64) (structs.User, error)
	GetUserByName(name string) (bool, error)
	CreateUserOauth(name string) error
}

// Post actions
type PostRedact interface {
	CreatePost(post *structs.Post) error
	GetUSerID(token string) (int64, error)
	GetUserName(userID int64) (string, error)
	GetPostBy(from, value string, user_id int64) (*structs.Post, error)
	LikePost(post *structs.Post) error
	DislikePost(post *structs.Post) error
	RedactContentPost(post *structs.Post) error
	DeletePost(post *structs.Post) error
	GetAllPosts() ([]structs.Post, error)
	GetAllLikedPosts(user_id int64) ([]structs.PostReaction, error)
	GetAllDislikedPosts(user_id int64) ([]structs.PostReaction, error)
	GetAllUserPosts(user_id int64) ([]structs.Post, error)
	GetFilteredPosts(java, kotlin, python, topic string) ([]structs.Post, error)
}

type CommentRedact interface {
	CreateComment(comm *structs.Comment) error
	GetAllComments(postID, userID int64) ([]structs.Comment, error)
	GetCommentByID(commentID int64) (structs.Comment, error)
	GetAllUserComments(userID int64) ([]structs.Comment, error)
	DeleteComment(comment structs.Comment) error
	UpdateCommentContent(comment structs.Comment) error
}

type PostReaction interface {
	LikePost(post_id, user_id, value int64) error
	FindReation(post_id, user_id, value int64) (*structs.PostReaction, error)
	CreateReaction(post_id, user_id, value int64) error
	DeleteReaction(post_id, user_id int64) error
	AllReactions(post_id int64) (int64, int64, error)
	GetPostReaction(user_id, post_id int64) (int64, error)
}

type CommentReaction interface {
	LikeComment(comment_id, user_id, value int64) error
	FindReation(comment_id, user_id, value int64) (*structs.PostReaction, error)
	CreateReaction(comment_id, user_id, value int64) error
	DeleteReaction(comment_id, user_id int64) error
	AllReactions(comment_id int64) (int64, int64, error)
	GetCommentReaction(user_id, commentId int64) (int64, error)
}

type NotificationPost interface {
	NotifyLikePost(post_id, user_id, author_id, value int64) error
	CreateNotifyReaction(post_id, user_id, author_id, value int64, username string) error
	DeletenNotifyReaction(post_id, user_id, author_id, value int64) error
	GetPostNotification(author_id int64) ([]structs.Notify, error)
}

type Repository struct {
	Authorization
	PostRedact
	CommentRedact
	PostReaction
	CommentReaction
	NotificationPost
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:    NewAuth(db),
		PostRedact:       NewPostRedactDB(db),
		CommentRedact:    NewCommentRedactDB(db),
		PostReaction:     NewReactionDB(db),
		CommentReaction:  NewCommentReactionDB(db),
		NotificationPost: NewNotifyDB(db),
	}
}
