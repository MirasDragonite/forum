package structs

type Post struct {
	Id             int64
	PostAuthorID   int64
	PostAuthorName string
	Topic          string
	Title          string
	Content        string
	Like           int64
	Dislike        int64
	Comments       []Comment
}

func CreatePost(id int64, authorID int64, topic string, title string, content string) *Post {
	return &Post{
		Id:           id,
		PostAuthorID: authorID,
		Topic:        topic,
		Title:        title,
		Content:      content,
		Like:         0,
		Dislike:      0,
		Comments:     []Comment{},
	}
}

// func (post *Post) ChangeContent(newContent string) {
// 	post.Content = newContent
// }

// func (post *Post) ChangeTitle(newTitle string) {
// 	post.Title = newTitle
// }

// func (post *Post) LikePost() {
// 	post.Like++
// 	post.Dislike--
// }

// func (post *Post) DislikePost() {
// 	post.Like--
// 	post.Dislike++
// }

// func (post *Post) WriteComment(commentID int64, CommentAuthorID int64, PostID int64, Content string) {
// 	post.Comments = append(post.Comments, *CreateComment(commentID, CommentAuthorID, PostID, Content))
// }

// func DeletePost(postID int64, authorID int64, user *User) {
// }
