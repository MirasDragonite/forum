package repository

import (
	"database/sql"
	"fmt"

	"forum/structs"
)

type CommentReactionDB struct {
	db *sql.DB
}

func NewCommentReactionDB(db *sql.DB) *CommentReactionDB {
	return &CommentReactionDB{db: db}
}

func (r *CommentReactionDB) LikeComment(comment_id, user_id, value int64) error {
	query := `UPDATE comment_reactions SET reaction=$1  WHERE comment_id=$2 AND user_id=$3`

	_, err := r.db.Exec(query, &value, &comment_id, &user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentReactionDB) AllReactions(comment_id int64) (int64, int64, error) {
	query := `SELECT COUNT(*) FROM comment_reactions WHERE comment_id=$1 AND reaction=1 `
	fmt.Println("Calling...")
	row := r.db.QueryRow(query, comment_id)
	var likes int64
	err := row.Scan(&likes)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, err
	}

	query = `SELECT COUNT(*) FROM comment_reactions WHERE comment_id=$1 AND reaction=-1 `
	fmt.Println("Calling...")
	row = r.db.QueryRow(query, comment_id)
	var dislikes int64
	err = row.Scan(&dislikes)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, err
	}

	return likes, dislikes, nil
}

func (r *CommentReactionDB) FindReation(comment_id, user_id, value int64) (*structs.PostReaction, error) {
	query := `SELECT * FROM comment_reactions WHERE comment_id=$1 AND user_id=$2`

	row := r.db.QueryRow(query, comment_id, user_id)
	var postReaction structs.PostReaction
	err := row.Scan(&postReaction.ID, &postReaction.PostID, &postReaction.UserID, &postReaction.Value)
	if err != nil {
		return nil, err
	}

	return &postReaction, nil
}

func (r *CommentReactionDB) CreateReaction(comment_id, user_id, value int64) error {
	query := `INSERT INTO comment_reactions(comment_id,user_id,reaction) VALUES($1,$2,$3)`

	_, err := r.db.Exec(query, &comment_id, &user_id, &value)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (r *CommentReactionDB) DeleteReaction(comment_id, user_id int64) error {
	query := `DELETE FROM comment_reactions WHERE comment_id=$1 AND user_id=$2`

	_, err := r.db.Exec(query, comment_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
