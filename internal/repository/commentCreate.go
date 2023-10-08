package repository

import (
	"database/sql"
	"errors"
	"forum/structs"
)

type CommentRedactDB struct {
	db *sql.DB
}

func NewCommentRedactDB(db *sql.DB) *CommentRedactDB {
	return &CommentRedactDB{db: db}
}

func (comm *CommentRedactDB) CreateComment(comment *structs.Comment) error {
	query := `INSERT INTO comments(comment_author_id, post_id, content, like, dislike) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	result, err := comm.db.Exec(query, comment.CommentAuthorID, comment.PostID, comment.Content, comment.Like, comment.Dislike)
	if err != nil {
		return errors.New("Error: cannot create new comment, Check CreateComment()")
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	comment.CommentID = lastID
	return nil
}
