package queue

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"queue/db"
)

type repo struct{}

func newRepo() *repo {
	return &repo{}
}

func (r *repo) getAll() ([]queue, error) {
	query := "select id, name_rus, name_kaz, next_free_slot_number, responsible_user_username from queue"
	rows, err := db.Db().Query(query)
	if err != nil {
		return nil, err
	}

	var queues []queue
	for rows.Next() {
		var q queue
		err = rows.Scan(&q.Id, &q.NameRus, &q.NameKaz, &q.NextFreeSlotNumber, &q.ResponsibleUserUsername)
		if err != nil {
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

func (r *repo) getById(id int) (*queue, error) {
	query := `select 
				id, 
				name_rus, 
				name_kaz, 
				next_free_slot_number, 
				current_slot_number,
				responsible_user_username 
			  from queue
			  where id = $1`
	row := db.Db().QueryRow(query, id)
	var q queue
	err := row.Scan(
		&q.Id,
		&q.NameKaz,
		&q.NameKaz,
		&q.NextFreeSlotNumber,
		&q.CurrentSlotNumber,
		&q.ResponsibleUserUsername,
	)
	return &q, err
}

func (r *repo) create(tx *sql.Tx, q *queue) error {
	query := `insert into queue (name_rus, name_kaz, responsible_user_username)
			  values ($1, $2, $3)`
	_, err := tx.Exec(query, q.NameRus, q.NameKaz, q.ResponsibleUserUsername)
	return err
}

func (r *repo) existsByUsername(username string) (bool, error) {
	return r.existsBy("responsible_user_username", username)
}

func (r *repo) existsByNameRus(nameRus string) (bool, error) {
	return r.existsBy("name_rus", nameRus)
}

func (r *repo) existsByNameKaz(nameKaz string) (bool, error) {
	return r.existsBy("name_kaz", nameKaz)
}

func (r *repo) existsBy(column, value string) (bool, error) {
	query := fmt.Sprintf(`select exists(select id
              							from queue
              							where %s = $1)`, column)

	row := db.Db().QueryRow(query, value)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

var errNoQueueDeleted = errors.New("no queue was deleted")

func (r *repo) deleteById(tx *sql.Tx, id int) (string, error) {
	query := `delete
			  from queue
			  where id = $1
			  returning responsible_user_username`
	row := tx.QueryRow(query, id)
	var username string
	err := row.Scan(&username)
	if err == sql.ErrNoRows {
		return "", errNoQueueDeleted
	}
	return username, err
}

func (r *repo) incrementNextFreeSlot(tx *sql.Tx, queueId int) error {
	query := `update queue
			  set next_free_slot_number = (select next_free_slot_number + 1
										   from queue
										   where id = $1)
			  where id = $1`
	_, err := tx.Exec(query, queueId)
	return err
}
