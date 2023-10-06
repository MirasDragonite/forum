package structs

type Post struct {
	id       int64
	authorID int64
	topic    string
	title    string
	content  string
	like     int64
	dislike  int64
	comments []Comment
}

func CreatePost(id int64, authorID int64, topic string, title string, content string) *Post {
	return &Post{
		id:       id,
		authorID: authorID,
		topic:    topic,
		title:    title,
		content:  content,
		like:     0,
		dislike:  0,
		comments: []Comment{},
	}
}
// func DeletePost(){

// }

func (post *Post) GetPostID() int64 {
	return post.id
}

func (post *Post) GetPostAuthorID() int64 {
	return post.authorID
}

func (post *Post) GetPostTopic() string {
	return post.topic
}

func (post *Post) GetPostTitle() string {
	return post.title
}

func (post *Post) GetPostContent() string {
	return post.content
}

func (post *Post) GetPostLikeNumber() int64 {
	return post.like
}

func (post *Post) GetPostDislikeNumber() int64 {
	return post.dislike
}

func (post *Post) GetPostComments() []Comment {
	return post.comments
}

// Methods

func (post *Post) ChangeContent(newContent string) {
	post.content = newContent
}

// func (post *Post) LikePost(user *User) {
// 	if user.GetUserID() == int64(post.authorID) {

// 	}
// }
