package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Transaction Roolback
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}

	// _, err = db.Exec("PRAGMA foreign_keys = ON;")
	// if err != nil {
	// 	return nil, err
	// }

	query := `DROP TABLE IF EXISTS users;
	CREATE TABLE users(id INTEGER PRIMARY KEY, username TEXT NOT NULL, email TEXT NOT NULL UNIQUE,hash_password TEXT NOT NULL,createdDate TEXT NOT NULL);
	DROP TABLE IF EXISTS tokens;
	CREATE TABLE tokens(id INTEGER PRIMARY KEY, user_id INTEGER,token TEXT NOT NULL,expaired_data TEXT NOT NULL, FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE);
	DROP TABLE IF EXISTS posts;
	CREATE TABLE posts(id INTEGER PRIMARY KEY NOT NULL, postAuthorID INTEGER, topic TEXT, title TEXT NOT NULL, content TEXT NOT NULL, like INTEGER, dislike INTEGER, username TEXT, FOREIGN KEY(postAuthorID) REFERENCES  users(id) ON DELETE CASCADE);
	DROP TABLE IF EXISTS comments;
	CREATE TABLE comments(id INTEGER PRIMARY KEY NOT NULL, comment_author_id INTEGER,commentAuthorName TEXT NOT NULL, post_id INTEGER, content TEXT NOT NULL, like INTEGER, dislike INTEGER, FOREIGN KEY(comment_author_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE); 
	DROP TABLE IF EXISTS post_reactions;
	CREATE TABLE post_reactions (id INTEGER PRIMARY KEY NOT NULL, post_id INTEGER, user_ID INTEGER, reaction INTEGER,FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE);
	DROP TABLE IF EXISTS comment_reactions;
	CREATE TABLE comment_reactions (id INTEGER PRIMARY KEY NOT NULL, comment_id INTEGER, user_ID INTEGER, reaction INTEGER,FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(comment_id) REFERENCES comments(id) ON DELETE CASCADE);
	DROP TABLE IF EXISTS post_notification;
	CREATE TABLE post_notification(id INTEGER PRIMARY KEY NOT NULL, user_id INTEGER, author_id INTEGER,post_id INTEGER , reaction INTEGER,username TEXT, FOREIGN KEY(user_id) REFERENCES users(id),FOREIGN KEY(author_id) REFERENCES users(id),FOREIGN KEY(post_id) REFERENCES posts(id));
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfuly connect to database")
	return db, nil
}
