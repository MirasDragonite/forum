package repository

import (
	"database/sql"
	"errors"
	"forum/structs"
)

type PostRedactDB struct {
	db *sql.DB
}

func NewPostRedactDB(db *sql.DB) *PostRedactDB {
	return &PostRedactDB{db: db}
}

func (pr *PostRedactDB) CreatePost(post *structs.Post) error {
	query := `INSERT INTO posts(postAuthorID, topic, title, content, like, dislike) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := pr.db.Exec(query, post.PostAuthorID, post.Topic, post.Title, post.Content, post.Like, post.Dislike)
	if err != nil {
		return errors.New("Error: cannot create new post, Check CreatePost()")
	}
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
