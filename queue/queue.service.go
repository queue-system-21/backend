package queue

import (
	"queue/db"
	"queue/user"
)

type service struct {
	repo *repo
}

func newService() *service {
	return &service{
		repo: newRepo(),
	}
}

func (s *service) getAll() ([]queue, error) {
	return s.repo.getAll()
}

func (s *service) create(name, responsibleUserUsername string) error {
	tx, err := db.Db().Begin()
	if err != nil {
		return err
	}

	if err = s.repo.create(tx, name, responsibleUserUsername); err != nil {
		tx.Rollback()
		return err
	}

	if err = user.SetRole(tx, responsibleUserUsername, "receptionist"); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *service) existsByUsername(username string) (bool, error) {
	return s.repo.existsByUsername(username)
}

type errNoQueueDeleted struct{}

func (n errNoQueueDeleted) Error() string {
	return "no queue was deleted"
}

func deleteById(id int) error {
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
