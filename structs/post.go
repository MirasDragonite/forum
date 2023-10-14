package structs

type Post struct {
	Id             int64  `json:"Id"`
	PostAuthorID   int64  `json:"PostAuthorID"`
	PostAuthorName string `json:"PostAuthorName"`
	Topic          string `json:"Topic"`
	Title          string `json:"Title"`
	Content        string `json:"Content"`
	Like           int64  `json:"Like"`
	Dislike        int64  `json:"Dislike"`
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
