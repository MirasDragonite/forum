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

	query := `CREATE TABLE IF NOT  EXISTS users(id INTEGER PRIMARY KEY, username TEXT NOT NULL, email TEXT NOT NULL UNIQUE,hash_password TEXT NOT NULL,createdDate TEXT NOT NULL);
	
	CREATE TABLE IF NOT  EXISTS tokens(id INTEGER PRIMARY KEY, user_id INTEGER,token TEXT NOT NULL,expaired_data TEXT NOT NULL, FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE);
	
	CREATE TABLE IF NOT  EXISTS posts(id INTEGER PRIMARY KEY NOT NULL, postAuthorID INTEGER, topic TEXT, title TEXT NOT NULL, content TEXT NOT NULL, like INTEGER, dislike INTEGER, username TEXT, FOREIGN KEY(postAuthorID) REFERENCES  users(id) ON DELETE CASCADE);
	
	CREATE TABLE IF NOT  EXISTS comments(id INTEGER PRIMARY KEY NOT NULL, comment_author_id INTEGER,commentAuthorName TEXT NOT NULL, post_id INTEGER, content TEXT NOT NULL, like INTEGER, dislike INTEGER, FOREIGN KEY(comment_author_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE); 
	
	CREATE TABLE IF NOT  EXISTS post_reactions (id INTEGER PRIMARY KEY NOT NULL, post_id INTEGER, user_ID INTEGER, reaction INTEGER,FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE);
	
	CREATE TABLE IF NOT  EXISTS comment_reactions (id INTEGER PRIMARY KEY NOT NULL, comment_id INTEGER, user_ID INTEGER, reaction INTEGER,FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(comment_id) REFERENCES comments(id) ON DELETE CASCADE);
	
	CREATE TABLE IF NOT  EXISTS post_notification(id INTEGER PRIMARY KEY NOT NULL, user_id INTEGER, author_id INTEGER,post_id INTEGER , reaction INTEGER,username TEXT, FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE, FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfuly connect to database")
	return db, nil
}
