package user

import (
	"database/sql"
	"queue/db"
)

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) save(user *User) error {
	query := "insert into \"user\" (username, password) values ($1, $2);"
	_, err := db.Db().Exec(query, user.Username, user.Password)
	return err
}

func (r *repo) exists(user *User) (bool, error) {
	query := `select exists(select *
              from "user"
              where username = $1
                and password = $2);`
	row := db.Db().QueryRow(query, user.Username, user.Password)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

func (r *repo) updateRole(tx *sql.Tx, user *User) error {
	query := `update "user"
			set role_code = $2
			where username = $1`
	_, err := tx.Exec(query, user.Username, user.RoleCode)
	return err
}

func (r *repo) getRoleByUsername(username string) (string, error) {
	query := `select role_code
			  from "user"
			  where username = $1`
	var role string
	row := db.Db().QueryRow(query, username)
	err := row.Scan(&role)
	return role, err
}
