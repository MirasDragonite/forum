package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}

	query := `DROP TABLE IF EXISTS users;
	CREATE TABLE users(id INTEGER PRIMARY KEY, username TEXT NOT NULL UNIQUE, email TEXT NOT NULL UNIQUE,hash_password TEXT NOT NULL);
	DROP TABLE IF EXISTS tokens;
	CREATE TABLE tokens(id INTEGER PRIMARY KEY, user_id INTEGER,token TEXT NOT NULL,expaired_data TEXT NOT NULL, FOREIGN KEY(user_id) REFERENCES users(id));
	
	DROP TABLE IF EXISTS posts;
	CREATE TABLE posts(id INTEGER PRIMARY KEY NOT NULL, postAuthorID INTEGER, topic TEXT, title TEXT, content TEXT, like INTEGER, dislike INTEGER, FOREIGN KEY(postAuthorID) REFERENCES  users(id));
	
	DROP TABLE IF EXISTS comments;
	CREATE TABLE comments(id INTEGER PRIMARY KEY NOT NULL, comment_author_id INTEGER, post_id INTEGER, content TEXT, like INTEGER, dislike INTEGER, FOREIGN KEY(comment_author_id) REFERENCES users(id), FOREIGN KEY(post_id) REFERENCES posts(id));
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfuly connect to database")
	return db, nil
}