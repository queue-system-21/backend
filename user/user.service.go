package user

import (
	"database/sql"
	"queue/db"
)

func Create(username, password string) error {
	query := "insert into \"user\" (username, password) values ($1, $2);"
	_, err := db.Db().Exec(query, username, password)
	return err
}

func ValidateCredentials(username, password string) bool {
	query := `select exists(select *
              from "user"
              where username = $1
                and password = $2);`
	row := db.Db().QueryRow(query, username, password)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func SetRole(tx *sql.Tx, username, code string) error {
	query := `update "user"
			set role_code = $2
			where username = $1`
	_, err := tx.Exec(query, username, code)
	return err
}
