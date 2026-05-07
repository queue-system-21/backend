package queue

import "queue/db"

type userQueueNumberRepo struct{}

func newUserQueueNumberRepo() *userQueueNumberRepo {
	return &userQueueNumberRepo{}
}

func (r *userQueueNumberRepo) save(uqn *userQueueNumber) error {
	query := "insert into user_queue_number (username, queue_id, number) values ($1, $2, $3)"
	_, err := db.Db().Exec(query, uqn.Username, uqn.QueueId, uqn.Number)
	return err
}

func (r *userQueueNumberRepo) getByUsername(username string) (*userQueueNumber, error) {
	query := "select id, username, queue_id, number from user_queue_number where username = $1"
	var uqn userQueueNumber
	row := db.Db().QueryRow(query, username)
	err := row.Scan(&uqn.Id, &uqn.Username, &uqn.QueueId, &uqn.Number)
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
