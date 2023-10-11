package structs

type Comment struct {
	CommentID         int64  `json:"id"`
	CommentAuthorID   int64  `json:"author_id"`
	CommentAuthorName string `json:"authorName"`
	PostID            int64  `json:"post_id"`
	Content           string `json:"content"`
	Like              int64  `json:"like"`
	Dislike           int64  `json:"dislike"`
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
