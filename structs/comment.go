package structs

type Comment struct {
	CommentID       int64
	CommentAuthorID int64
	PostID          int64
	Content         string
	Like            int64
	Dislike         int64
}

func CreateComment(commentID int64, CommentAuthorID int64, PostID int64, Content string) *Comment {
	return &Comment{
		CommentID:       commentID,
		CommentAuthorID: CommentAuthorID,
		PostID:          PostID,
		Content:         Content,
		Like:            0,
		Dislike:         0,
	}
}

func (comment *Comment) LikeComment() {
	comment.Like++
	comment.Dislike--
}

func (comment *Comment) DislikeComment() {
	comment.Like--
	comment.Dislike++
}


func DeleteComment(post *Post, commentID int64) *Comment {
	for _, comment := range post.Comments {
		if comment.CommentID == commentID {
			return &comment
		}
	}
	return nil
}