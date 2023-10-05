package repository

import "database/sql"

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SOME QUERY`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	return db, nil
}
