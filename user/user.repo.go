package user

import (
	"database/sql"
	"queue/db"
)

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) create(username, password string) error {
	query := "insert into \"user\" (username, password) values ($1, $2);"
	_, err := db.Db().Exec(query, username, password)
	return err
}

func (r *repo) exists(username, password string) (bool, error) {
	query := `select exists(select *
              from "user"
              where username = $1
                and password = $2);`
	row := db.Db().QueryRow(query, username, password)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *repo) setRole(tx *sql.Tx, username, code string) error {
	query := `update "user"
			set role_code = $2
			where username = $1`
	_, err := tx.Exec(query, username, code)
	return err
}

func (r *repo) getRole(username, password string) (string, error) {
	query := `select role_code
			  from "user"
			  where username = $1
				and password = $2`
	var role string
	row := db.Db().QueryRow(query, username, password)
	err := row.Scan(&role)
	return role, err
}
