package repository

import (
	"database/sql"
	"fmt"

	"forum/structs"
)

type ReactionDB struct {
	db *sql.DB
}

func NewReactionDB(db *sql.DB) *ReactionDB {
	return &ReactionDB{db: db}
}

func (r *ReactionDB) LikePost(post_id, user_id, value int64) error {
	fmt.Println("Here2")
	query := `UPDATE post_reactions SET reaction=$1  WHERE post_id=$2 AND user_id=$3`

	_, err := r.db.Exec(query, &value, &post_id, &user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ReactionDB) AllReactions(post_id int64) (int64, int64, error) {
	query := `SELECT COUNT(*) FROM post_reactions WHERE post_id=$1 AND reaction=1 `

	row := r.db.QueryRow(query, &post_id)
	var likes int64
	err := row.Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	query = `SELECT COUNT(*) FROM post_reactions WHERE post_id=$1 AND reaction=-1 `

	row = r.db.QueryRow(query, &post_id)
	var dislikes int64
	err = row.Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}

func (r *ReactionDB) FindReation(post_id, user_id, value int64) (*structs.PostReaction, error) {
	fmt.Println("Here1")
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
	fmt.Println("Here3")
	query := `INSERT INTO post_reactions(post_id,user_id,reaction) VALUES($1,$2,$3)`

	_, err := r.db.Exec(query, &post_id, &user_id, &value)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (r *ReactionDB) DeleteReaction(post_id, user_id int64) error {
	query := `DELETE FROM post_reactions WHERE post_id=$1 AND user_id=$2`

	_, err := r.db.Exec(query, post_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
