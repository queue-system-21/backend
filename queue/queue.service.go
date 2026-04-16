package queue

import (
	"queue/db"
	"queue/user"
)

type service struct {
	repo        *repo
	userService *user.Service
}

func newService() *service {
	return &service{
		repo:        newRepo(),
		userService: user.NewService(),
	}
}

func (s *service) getAll() ([]queue, error) {
	return s.repo.getAll()
}

func (s *service) create(nameRus, nameKaz, responsibleUserUsername string) error {
	tx, err := db.Db().Begin()
	if err != nil {
		return err
	}

	err = s.repo.create(tx, nameRus, nameKaz, responsibleUserUsername)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.userService.SetRole(tx, responsibleUserUsername, "receptionist")
	if err != nil {
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
