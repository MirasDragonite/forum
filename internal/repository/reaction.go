package repository

import (
	"database/sql"

	"forum/structs"
)

type ReactionDB struct {
	db *sql.DB
}

func NewReactionDB(db *sql.DB) *ReactionDB {
	return &ReactionDB{db: db}
}

func (r *ReactionDB) LikePost(post_id, user_id, value int64) error {
	query := `UPDATE post_reactions SET reaction=$1  WHERE post_id=$2 AND user_id=$3`

	_, err := r.db.Exec(query, &value, &post_id, &user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReactionDB) FindReation(post_id, user_id, value int64) (*structs.PostReaction, error) {
	query := `SELECT * FROM post_reactions WHERE post_id=$1 AND user_id=$2`

	row := r.db.QueryRow(query, post_id, user_id)
	var postReaction structs.PostReaction
	err := row.Scan(&postReaction.ID, &postReaction.PostID, &postReaction.UserID, &postReaction.Value)
	if err != nil {
		return nil, err
	}

	return &postReaction, nil
}

func (r *ReactionDB) CreateReaction(post_id, user_id, value int64) error {
	query := `INSERT INTO post_reaction(post_id,user_id, reaction) VALUE($1,$2,$3)`

	_, err := r.db.Exec(query, &post_id, &user_id, &value)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReactionDB) DeleteReaction(post_id, user_id int64) error {
	query := `DELETE FROM post_reaction WHERE post_id=$1 AND user_id=$2`

	_, err := r.db.Exec(query, post_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
