package queue

import (
	"database/sql"
	"log"
	"queue/db"
)

type queue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) getAll() ([]queue, error) {
	query := "select id, name from queue"
	rows, err := db.Db().Query(query)
	if err != nil {
		return nil, err
	}

	var queues []queue
	for rows.Next() {
		var q queue
		if err = rows.Scan(&q.Id, &q.Name); err != nil {
			log.Println("Error in queue.getAll:", err)
			continue
		}
		queues = append(queues, q)
	}
	if queues == nil {
		queues = make([]queue, 0)
	}
	return queues, nil
}

func (r *repo) create(tx *sql.Tx, name, responsibleUserUsername string) error {
	query := `insert into queue (name, responsible_user_username) values ($1, $2)`
	_, err := tx.Exec(query, name, responsibleUserUsername)
	return err
}

func (r *repo) existsByUsername(username string) (bool, error) {
	query := `select exists(select id
              from queue
              where responsible_user_username = $1)`

	row := db.Db().QueryRow(query, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
