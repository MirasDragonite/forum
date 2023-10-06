package structs

type Post struct {
	id        int
	authorID int
	content   string
	like     int
	dislike   int
	comments  []Comment
}



func (post *Post) GetPostID() int {
	return post.id
}
func (post *Post) GetPostAuthorID() int {
	return post.authorID
}

func (post *Post) GetPostContent() string {
	return post.content
}

func (post *Post) GetPostLikeNumber() int {
	return post.like
}

func (post *Post) GetPostDislikeNumber() int {
	return post.dislike
}

func (post *Post) GetPostComments() []Comment {
	return post.comments
} 


// Methods 

func (post *Post) ChangeContent(newContent string) {
	post.content = newContent
}

func (post *Post) LikePost(user *User) {
	if user.
}	