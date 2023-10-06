package structs

type Comment struct {
	author_id int
	content   string
	like      int
	dislike   int
}

func (comment *Comment) getCommentAuthorID() int {
	return comment.author_id
}

func (comment *Comment) getCommentContent() string {
	return comment.content
}

func (comment *Comment) getCommentLikeNumber() int {
	return comment.like
}

func (comment *Comment) getCommentDislikeNumber() int {
	return comment.dislike
}
