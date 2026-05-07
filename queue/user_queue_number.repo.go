package queue

import (
	"database/sql"
	"queue/db"
)

type userQueueNumberRepo struct{}

func newUserQueueNumberRepo() *userQueueNumberRepo {
	return &userQueueNumberRepo{}
}

func (r *userQueueNumberRepo) save(tx *sql.Tx, uqn *userQueueNumber) error {
	query := "insert into user_queue_number (username, queue_id, number) values ($1, $2, $3)"
	_, err := tx.Exec(query, uqn.Username, uqn.QueueId, uqn.Number)
	return err
}

func (r *userQueueNumberRepo) getByUsername(username string) (*userQueueNumber, error) {
	query := `select 
				uqn.id, 
				uqn.username, 
				uqn.queue_id, 
				uqn.number - q.current_slot_number, 
				q.name_rus, 
				q.name_kaz
			  from user_queue_number uqn
				left join queue q on uqn.queue_id = q.id
			  where uqn.username = $1`
	var uqn userQueueNumber
	row := db.Db().QueryRow(query, username)
	err := row.Scan(
		&uqn.Id,
		&uqn.Username,
		&uqn.QueueId,
		&uqn.Number,
		&uqn.queue.NameRus,
		&uqn.queue.NameKaz,
	)
	return &uqn, err
}

func (r *userQueueNumberRepo) existsByUsername(username string) (bool, error) {
	query := "select exists(select id from user_queue_number where username = $1)"
	var exists bool
	row := db.Db().QueryRow(query, username)
	err := row.Scan(&exists)
	return exists, err
}

func (r *userQueueNumberRepo) deleteByUsername(username string) error {
	query := "delete from user_queue_number where username = $1"
	_, err := db.Db().Exec(query, username)
	return err
}
