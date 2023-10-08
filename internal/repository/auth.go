package repository

import (
	"database/sql"
	"fmt"

	"forum/structs"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user *structs.User) (int64, error) {
	query := `INSERT INTO users(username,email,hash_password) VALUES($1,$2,$3) RETURNING id`

	result, err := r.db.Exec(query, user.Username, user.Email, user.HashedPassword)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Auth) GetUserBy(element, from string) (structs.User, error) {
	var user structs.User

	query := fmt.Sprintf("SELECT * FROM users WHERE %s=$1 ", from)
	row := r.db.QueryRow(query, element)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword)
	if err != nil {
		return structs.User{}, err
	}
	return user, nil
}

func (r *Auth) CreateToken(user structs.User, token, expaired_data string) error {
	query := `INSERT INTO tokens(user_id,token,expaired_data ) VALUES($1,$2,$3)`

	_, err := r.db.Exec(query, &user.Id, &token, &expaired_data)
	if err != nil {
		return err
	}

	return nil
}

func (r *Auth) GetSession(userID int64) (structs.Session, error) {
	var session structs.Session

	query := `SELECT * FROM tokens WHERE user_id=$1`

	row := r.db.QueryRow(query, userID)
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpairedData)
	if err != nil {
		return structs.Session{}, err
	}
	return session, nil
}

func (r *Auth) UpdateToken(user structs.User, token, expaired_data string) error {
	query := `UPDATE tokens SET token=$1 ,expaired_data=$2 WHERE user_id=$3`

	_, err := r.db.Exec(query, &token, &expaired_data, &user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Auth) DeleteToken(token string) error {
	query := `DELETE FROM tokens WHERE token=$1`

	_, err := r.db.Exec(query, &token)
	if err != nil {
		return err
	}
	return nil
}
