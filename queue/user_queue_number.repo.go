package queue

import "queue/db"

type userQueueNumberRepo struct{}

func newUserQueueNumberRepo() *userQueueNumberRepo {
	return &userQueueNumberRepo{}
}

func (r *userQueueNumberRepo) save(uqn *userQueueNumber) error {
	query := "insert into user_queue_number (username, queue_id) values ($1, $2)"
	_, err := db.Db().Exec(query, uqn.Username, uqn.QueueId)
	return err
}
