package queue

import (
	"errors"
	"queue/db"
	"queue/user"
)

var errUserBusy = errors.New("You cannot assign this user for this queue")
var errNameRusNotUnique = errors.New("nameRus is not unique")
var errNameKazNotUnique = errors.New("nameKaz is not unique")

type service struct {
	repo                *repo
	userQueueNumberRepo *userQueueNumberRepo
	userService         *user.Service
}

func newService() *service {
	return &service{
		repo:                newRepo(),
		userQueueNumberRepo: newUserQueueNumberRepo(),
		userService:         user.NewService(),
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

	q := queue{
		NameRus:                 nameRus,
		NameKaz:                 nameKaz,
		ResponsibleUserUsername: responsibleUserUsername,
	}
	if err = s.repo.create(tx, &q); err != nil {
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

func (s *service) existsBy(username, nameRus, nameKaz string) error {
	exists, err := s.repo.existsByUsername(username)
	if err != nil {
		return err
	}
	if exists {
		return errUserBusy
	}

	exists, err = s.repo.existsByNameRus(nameRus)
	if err != nil {
		return err
	}
	if exists {
		return errNameRusNotUnique
	}

	exists, err = s.repo.existsByNameKaz(nameKaz)
	if err != nil {
		return err
	}
	if exists {
		return errNameKazNotUnique
	}

	return nil
}

func (s *service) deleteById(id int) error {
	tx, err := db.Db().Begin()
	if err != nil {
		return err
	}

	username, err := s.repo.deleteById(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = s.userService.SetRole(tx, username, "user"); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *service) getUserRole(username string) (string, error) {
	return s.userService.GetRoleByUsername(username)
}

func (s *service) join(username string, queueId int) error {
	uqn := userQueueNumber{
		Username: username,
		QueueId:  queueId,
	}
	return s.userQueueNumberRepo.save(&uqn)
}

func (s *service) getQueueNumber(username string) (int, error) {
	return s.userQueueNumberRepo.getNumber(username)
}
