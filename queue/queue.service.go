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

func (s *service) deleteById(id int) error {
	return s.repo.deleteById(id)
}
