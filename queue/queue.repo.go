package queue

import (
	"database/sql"
	"log"
	"queue/db"
)

type queue struct {
	Id      int    `json:"id"`
	NameRus string `json:"nameRus"`
	NameKaz string `json:"nameKaz"`
}

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) getAll() ([]queue, error) {
	query := "select id, name_rus, name_kaz from queue"
	rows, err := db.Db().Query(query)
	if err != nil {
		return nil, err
	}

	var queues []queue
	for rows.Next() {
		var q queue
		if err = rows.Scan(&q.Id, &q.NameRus, &q.NameKaz); err != nil {
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

func (r *repo) create(tx *sql.Tx, nameRus, nameKaz, responsibleUserUsername string) error {
	query := `insert into queue (name_rus, name_kaz, responsible_user_username)
			  values ($1, $2, $3)`
	_, err := tx.Exec(query, nameRus, nameKaz, responsibleUserUsername)
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

type errNoQueueDeleted struct{}

func (n errNoQueueDeleted) Error() string {
	return "no queue was deleted"
}

func (r *repo) deleteById(id int) error {
	query := "delete from queue where id = $1"
	res, err := db.Db().Exec(query, id)
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return errNoQueueDeleted{}
	}
	return err
}
