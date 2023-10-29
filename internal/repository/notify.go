package repository

import (
	"database/sql"
)

type NotifyDB struct {
	db *sql.DB
}

func NewNotifyDB(db *sql.DB) *PostReactionDB {
	return &PostReactionDB{db: db}
}

func (r *PostReactionDB) NotifyLikePost(post_id, user_id, value int64) error {
	query := `UPDATE post_reactions SET reaction=$1  WHERE post_id=$2 AND user_id=$3`

	_, err := r.db.Exec(query, &value, &post_id, &user_id)
	if err != nil {
		return err
	}

	return nil
}

// func (r *PostReactionDB) AllReactions(post_id int64) (int64, int64, error) {
// 	query := `SELECT COUNT(*) FROM post_reactions WHERE post_id=$1 AND reaction=1 `

// 	row := r.db.QueryRow(query, &post_id)
// 	var likes int64
// 	err := row.Scan(&likes)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	query = `SELECT COUNT(*) FROM post_reactions WHERE post_id=$1 AND reaction=-1 `

// 	row = r.db.QueryRow(query, &post_id)
// 	var dislikes int64
// 	err = row.Scan(&dislikes)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	return likes, dislikes, nil
// }

// func (r *PostReactionDB) FindReation(post_id, user_id, value int64) (*structs.PostReaction, error) {
// 	query := `SELECT * FROM post_reactions WHERE post_id=$1 AND user_id=$2`

// 	row := r.db.QueryRow(query, post_id, user_id)
// 	var postReaction structs.PostReaction
// 	err := row.Scan(&postReaction.ID, &postReaction.PostID, &postReaction.UserID, &postReaction.Value)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &postReaction, nil
// }

// func (r *PostReactionDB) CreateReaction(post_id, user_id, value int64) error {
// 	query := `INSERT INTO post_reactions(post_id,user_id,reaction) VALUES($1,$2,$3)`

// 	_, err := r.db.Exec(query, &post_id, &user_id, &value)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}

// 	return nil
// }

func (r *PostReactionDB) DeletenNotifyReaction(post_id, user_id int64) error {
	query := `DELETE FROM post_reactions WHERE post_id=$1 AND user_id=$2`

	_, err := r.db.Exec(query, post_id, user_id)
	if err != nil {
		return err
	}
	return nil
}

// func (r *PostReactionDB) GetPostReaction(user_id, post_id int64) (int64, error) {
// 	query := `SELECT id,reaction FROM post_reactions WHERE post_id=$1 AND user_ID=$2`

// 	var id, reaction int64
// 	row := r.db.QueryRow(query, post_id, user_id)
// 	err := row.Scan(&id, &reaction)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return reaction, nil
// }
