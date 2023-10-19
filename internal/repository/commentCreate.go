package repository

import (
	"database/sql"
	"errors"
	"fmt"
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

func (comm *CommentRedactDB) GetAllComments(postID, userID int64) ([]structs.Comment, error) {
	var comments []structs.Comment
	query := `SELECT * from comments WHERE post_id = $1`
	rows, err := comm.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var comment structs.Comment
		err := rows.Scan(&comment.CommentID, &comment.CommentAuthorID, &comment.PostID, &comment.Content, &comment.Like, &comment.Dislike)
		if err != nil {
			return nil, err
		}

		query := `SELECT id,reaction FROM comment_reactions WHERE comment_id=$1 AND user_ID=$2`
		var id, reaction int64
		fmt.Println("COmmentid:", comment.CommentID)
		fmt.Println("UserID:", userID)
		row := comm.db.QueryRow(query, comment.CommentID, userID)
		err = row.Scan(&id, &reaction)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows")
				comment.Liked = false
				comment.Disliked = false
			} else {
				return nil, err
			}
		}

		if reaction == 1 {
			fmt.Println("Comment Liked")
			comment.Liked = true
			comment.Disliked = false
		} else if reaction == -1 {
			comment.Liked = false
			comment.Disliked = true
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (comm *CommentRedactDB) GetCommentByID(commentID int64) (structs.Comment, error) {
	fmt.Println("cOMMentid", commentID)
	query := `SELECT * FROM comments WHERE id=$1`
	var comment structs.Comment
	row := comm.db.QueryRow(query, &commentID)
	err := row.Scan(&comment.CommentID, &comment.CommentAuthorID, &comment.PostID, &comment.Content, &comment.Like, &comment.Dislike)
	if err != nil {
		return structs.Comment{}, err
	}
	return comment, nil
}
