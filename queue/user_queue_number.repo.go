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

func (r *userQueueNumberRepo) getNumber(username string) (int, error) {
	query := `select num
			  from (select username, row_number() over () num
				    from user_queue_number
				    where queue_id = (select queue_id
									  from user_queue_number
									  where username = $1)
				    order by id) t
			  where username = $1`
	var num int
	row := db.Db().QueryRow(query, username)
	err := row.Scan(&num)
	return num, err
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
