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

func (r *Auth) CreateUser(user *structs.User) error {
	query := `INSERT INTO users(username,email,hash_password,createdDate) VALUES($1,$2,$3,$4) `

	_, err := r.db.Exec(query, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedDate)
	if err != nil {
		return err
	}
	//
	// id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (r *Auth) GetUserBy(element, from string) (structs.User, error) {
	var user structs.User

	query := fmt.Sprintf("SELECT * FROM users WHERE %s=$1 ", from)
	row := r.db.QueryRow(query, element)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedDate)
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

func (r *Auth) GetSessionByToken(token string) (structs.Session, error) {
	var session structs.Session

	query := `SELECT * FROM tokens WHERE token=$1`

	row := r.db.QueryRow(query, token)
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpairedData)
	if err != nil {
		return structs.Session{}, err
	}
	return session, nil
}

func (r *Auth) GetUserById(id int64) (structs.User, error) {
	var user structs.User

	query := `SELECT * FROM users WHERE id=$1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedDate)
	if err != nil {
		return structs.User{}, err
	}
	return user, nil
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

func (r *Auth) GetUserByName(name string) (bool, error) {
	var user structs.User
	query := `SELECT * FROM users WHERE username=$1`
	row := r.db.QueryRow(query, name)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedDate)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (r *Auth) CreateUserOauth(name string) error {
	query := `INSERT INTO users(username,email,hash_password,createdDate) VALUES($1,$2,$3,$4) `

	_, err := r.db.Exec(query, name, name, "", "")
	if err != nil {
		return err
	}
	//
	// id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
