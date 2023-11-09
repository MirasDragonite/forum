package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/structs"
	"strconv"
	"strings"
)

type PostRedactDB struct {
	db *sql.DB
}

func NewPostRedactDB(db *sql.DB) *PostRedactDB {
	return &PostRedactDB{db: db}
}

func (pr *PostRedactDB) CreatePost(post *structs.Post) error {
	query := `INSERT INTO posts(postAuthorID, topic, title, content, like, dislike, username) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result, err := pr.db.Exec(query, post.PostAuthorID, post.TopicString, post.Title, post.Content, post.Like, post.Dislike, post.PostAuthorName)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error: cannot create new post, Check CreatePost()")
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = lastID

	return nil
}

func (pr *PostRedactDB) GetUSerID(token string) (int64, error) {
	query := `SELECT user_id FROM tokens WHERE token=$1`
	row := pr.db.QueryRow(query, token)
	var userID int64
	err := row.Scan(&userID)
	if err != nil {
		return 0, errors.New("No such token in the db")
	}
	return userID, nil
}

func (pr *PostRedactDB) GetUserName(userID int64) (string, error) {
	query := `SELECT username FROM users WHERE id=$1`
	row := pr.db.QueryRow(query, userID)
	var userName string
	err := row.Scan(&userName)
	if err != nil {
		return "", errors.New("Error in func GetUserName() with scanning userName value")
	}
	return userName, nil
}

func (pr *PostRedactDB) GetPostBy(from, value string, user_id int64) (*structs.Post, error) {
	var post structs.Post
	if from == "id" || from == "postAuthorID" {
		value, err2 := strconv.Atoi(value)
		if err2 != nil {
			return &structs.Post{}, err2
		}
		query := fmt.Sprintf(`SELECT * FROM posts WHERE %s = $1`, from)

		row := pr.db.QueryRow(query, value)

		err := row.Scan(&post.Id, &post.PostAuthorID, &post.TopicString, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.PostAuthorName)
		if err != nil {
			return &structs.Post{}, err
		}
		var comments []structs.Comment
		query = `SELECT * from comments WHERE post_id = $1`
		rows, err := pr.db.Query(query, post.Id)
		if err != nil {
			return &structs.Post{}, err
		}
		defer rows.Close()
		post.Topic = strings.Split(post.TopicString, "|")
		for rows.Next() {
			var comment structs.Comment
			err := rows.Scan(&comment.CommentID, &comment.CommentAuthorID, &comment.CommentAuthorName, &comment.PostID, &comment.Content, &comment.Like, &comment.Dislike)
			if err != nil {
				return &structs.Post{}, err
			}

			comments = append(comments, comment)
		}
		post.Comments = comments
	} else {
		query := fmt.Sprintf(`SELECT * FROM posts WHERE %s = $1`, from)

		row := pr.db.QueryRow(query, value)
		err := row.Scan(&post.Id, &post.PostAuthorID, &post.TopicString, &post.Title, &post.Content, &post.Like, &post.Dislike)
		if err != nil {
			return &structs.Post{}, err
		}
		var comments []structs.Comment
		query = `SELECT * from comments WHERE post_id = $1`
		rows, err := pr.db.Query(query, post.Id)
		if err != nil {
			return &structs.Post{}, err
		}
		post.Topic = strings.Split(post.TopicString, "|")
		defer rows.Close()
		for rows.Next() {
			var comment structs.Comment
			err := rows.Scan(&comment.CommentID, &comment.CommentAuthorID, &comment.CommentAuthorName, &comment.PostID, &comment.Content, &comment.Like, &comment.Dislike)
			if err != nil {
				return &structs.Post{}, err
			}

			comments = append(comments, comment)
		}
		post.Comments = comments
	}

	if user_id != 0 {
		query := `SELECT id,reaction FROM post_reactions WHERE post_id=$1 AND user_ID=$2`

		var id, reaction int64
		row := pr.db.QueryRow(query, post.Id, user_id)
		err := row.Scan(&id, &reaction)
		if err != nil {
			if err == sql.ErrNoRows {
				post.Liked = false
				post.Disliked = false
			} else {
				return nil, err
			}
		}
		if reaction == 1 {
			post.Liked = true
			post.Disliked = false
		} else if reaction == -1 {
			post.Liked = false
			post.Disliked = true
		}
	}
	return &post, nil
}

func (pr *PostRedactDB) LikePost(post *structs.Post) error {
	query := `UPDATE posts SET like = $1 dislike = $2 WHERE id = $3;`
	_, err := pr.db.Exec(query, post.Like, post.Dislike, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRedactDB) DislikePost(post *structs.Post) error {
	query := `UPDATE posts SET like = $1 dislike = $2 WHERE id = $3;`
	_, err := pr.db.Exec(query, post.Like, post.Dislike, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRedactDB) RedactContentPost(post *structs.Post) error {
	query := `UPDATE posts SET content = $1 , title=$2 WHERE id = $3;`
	_, err := pr.db.Exec(query, post.Content, post.Title, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRedactDB) DeletePost(post *structs.Post) error {
	deletePostQuery := "DELETE FROM posts WHERE id = $1;"
	_, err := pr.db.Exec(deletePostQuery, post.Id)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRedactDB) GetAllPosts() ([]structs.Post, error) {
	query := `SELECT * FROM posts ORDER BY id DESC ;`

	var posts []structs.Post

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.Id, &post.PostAuthorID, &post.TopicString, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.PostAuthorName)
		if err != nil {
			return nil, err
		}
		post.Topic = strings.Split(post.TopicString, "|")
		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRedactDB) GetAllLikedPosts(user_id int64) ([]structs.PostReaction, error) {
	query := `SELECT * FROM post_reactions WHERE user_id=$1 AND reaction=1 ORDER BY id DESC`

	var posts []structs.PostReaction

	rows, err := pr.db.Query(query, &user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.PostReaction
		err := rows.Scan(&post.ID, &post.PostID, &post.UserID, &post.Value)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRedactDB) GetAllUserPosts(user_id int64) ([]structs.Post, error) {
	query := `SELECT * FROM posts WHERE postAuthorID=$1 ORDER BY id DESC`

	var posts []structs.Post

	rows, err := pr.db.Query(query, &user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.Id, &post.PostAuthorID, &post.TopicString, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.PostAuthorName)
		if err != nil {
			return nil, err
		}
		post.Topic = strings.Split(post.TopicString, "|")

		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRedactDB) GetFilteredPosts(java, kotlin, python, topic string) ([]structs.Post, error) {
	query := `SELECT * FROM posts  WHERE topic LIKE ? AND topic LIKE ? AND topic LIKE ? AND topic LIKE ? ORDER BY id DESC`

	var posts []structs.Post
	var rows *sql.Rows
	var err error

	if len(java+kotlin+python+strings.TrimSpace(topic)) == 0 {
		rows, err = pr.db.Query(query, "", "", "", "")
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = pr.db.Query(query, "%"+java+"%", "%"+kotlin+"%", "%"+python+"%", "%"+topic+"%")
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.Id, &post.PostAuthorID, &post.TopicString, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.PostAuthorName)
		if err != nil {
			return nil, err
		}
		post.Topic = strings.Split(post.TopicString, "|")
		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRedactDB) GetAllDislikedPosts(user_id int64) ([]structs.PostReaction, error) {
	query := `SELECT * FROM post_reactions WHERE user_id=$1 AND reaction=-1 ORDER BY id DESC`

	var posts []structs.PostReaction

	rows, err := pr.db.Query(query, &user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post structs.PostReaction
		err := rows.Scan(&post.ID, &post.PostID, &post.UserID, &post.Value)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
