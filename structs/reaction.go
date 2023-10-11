package structs

type PostReaction struct {
	ID int64
	AuthorID int64
	PostID  int64
	Like bool
	Dislike bool

}

type CommentReaction struct {
	ID int64
	AuthorID int64
	PostID  int64
	Like bool
	Dislike bool
}