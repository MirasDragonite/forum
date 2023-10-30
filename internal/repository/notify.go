package repository

import (
	"database/sql"
	"fmt"

	"forum/structs"
)

type NotifyDB struct {
	db *sql.DB
}

func NewNotifyDB(db *sql.DB) *PostReactionDB {
	return &PostReactionDB{db: db}
}

func (r *PostReactionDB) NotifyLikePost(post_id, user_id, author_id, value int64) error {
	query := `UPDATE post_notification  SET reaction=$1  WHERE post_id=$2 AND user_id=$3 AND author_id=$4`

	_, err := r.db.Exec(query, &value, &post_id, &user_id, &author_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostReactionDB) CreateNotifyReaction(post_id, user_id, author_id, value int64, username string) error {
	query := `INSERT INTO post_notification(user_id,author_id,post_id,reaction,username) VALUES($1,$2,$3,$4,$5)`
	fmt.Println("Here in repo")
	_, err := r.db.Exec(query, &user_id, &author_id, &post_id, &value, &username)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (r *PostReactionDB) DeletenNotifyReaction(post_id, user_id, author_id int64) error {
	query := `DELETE FROM post_notification WHERE post_id=$1 AND user_id=$2 AND author_id=$3`

	_, err := r.db.Exec(query, post_id, user_id, author_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostReactionDB) GetPostNotification(author_id int64) ([]structs.Notify, error) {
	query := `SELECT * FROM post_notification WHERE author_id=$1 ORDER BY id DESC`

	var postNotifications []structs.Notify
	row, err := r.db.Query(query, author_id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var postNotification structs.Notify

		err = row.Scan(&postNotification.ID, &postNotification.UserID, &postNotification.AuthorID, &postNotification.PostID, &postNotification.Reaction, &postNotification.Username)
		if err != nil {
			return nil, err
		}
		postNotifications = append(postNotifications, postNotification)
	}

	fmt.Println("Notifications:", postNotifications)
	return postNotifications, nil
}
