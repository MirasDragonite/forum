package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"forum/structs"
)

type PostRedactDB struct {
	db *sql.DB
}

func NewPostRedactDB(db *sql.DB) *PostRedactDB {
	return &PostRedactDB{db: db}
}

func (pr *PostRedactDB) CreatePost(post *structs.Post) error {
	query := `INSERT INTO posts(postAuthorID, topic, title, content, like, dislike, username) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result, err := pr.db.Exec(query, post.PostAuthorID, post.Topic, post.Title, post.Content, post.Like, post.Dislike, post.PostAuthorName)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error: cannot create new post, Check CreatePost()")
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = lastID
	fmt.Println("Repo:", post)
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

func (pr *PostRedactDB) GetPostBy(from, value string) (*structs.Post, error) {
	var post structs.Post
	if from == "id" || from == "postAuthorID" {
		value, err2 := strconv.Atoi(value)
		if err2 != nil {
			return &structs.Post{}, err2
		}
		query := fmt.Sprintf(`SELECT * FROM posts WHERE %s = $1`, from)

		row := pr.db.QueryRow(query, value)

		err := row.Scan(&post.Id, &post.PostAuthorID, &post.Topic, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.PostAuthorName)
		if err != nil {
			return &structs.Post{}, err
		}
		var comments []structs.Comment
		query = `SELECT * from comments WHERE post_id = $1`
		rows, err := pr.db.Query(query, post.Id)
		if err != nil {
			return &structs.Post{}, err
		}
		for rows.Next() {
			var comment structs.Comment
			err := rows.Scan(&comment.CommentID, &comment.CommentAuthorID, &comment.PostID, &comment.Content, &comment.Like, &comment.Dislike)
			if err != nil {
				return &structs.Post{}, err
			}
			comments = append(comments, comment)
		}
		post.Comments = comments
	} else {
		query := fmt.Sprintf(`SELECT * FROM posts WHERE %s = $1`, from)

		row := pr.db.QueryRow(query, value)
		err := row.Scan(&post.Id, &post.PostAuthorID, &post.Topic, &post.Title, &post.Content, &post.Like, &post.Dislike)
		if err != nil {
			return &structs.Post{}, err
		}
		var comments []structs.Comment
		query = `SELECT * from comments WHERE post_id = $1`
		rows, err := pr.db.Query(query, post.Id)
		if err != nil {
			return &structs.Post{}, err
		}
		for rows.Next() {
			var comment structs.Comment
			err := rows.Scan(&comment.CommentID, &comment.CommentAuthorID, &comment.PostID, &comment.Content, &comment.Like, &comment.Dislike)
			if err != nil {
				return &structs.Post{}, err
			}
			comments = append(comments, comment)
		}
		post.Comments = comments
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
	query := `UPDATE posts SET content = $1 WHERE id = $2;`
	_, err := pr.db.Exec(query, post.Content, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRedactDB) DeletePost(post *structs.Post) error {
	query := `DELETE FROM posts WHERE id = $1;`
	_, err := pr.db.Exec(query, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRedactDB) GetAllPosts() ([]structs.Post, error) {
	query := `SELECT * FROM posts;`

	var posts []structs.Post

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.Id, &post.PostAuthorID, &post.Topic, &post.Title, &post.Content, &post.Like, &post.Dislike, &post.PostAuthorName)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
