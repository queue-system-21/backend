package queue

import (
	"errors"
	"log"
	"queue/db"
	"queue/user"
)

var errUserBusy = errors.New("You cannot assign this user for this queue")
var errNameRusNotUnique = errors.New("nameRus is not unique")
var errNameKazNotUnique = errors.New("nameKaz is not unique")
var errUserJoinedQueue = errors.New("User have already joined a queue")
var errUserQueueCheck = errors.New("Failed to check if the user have already joined a queue")
var errNextFreeCheck = errors.New("Failed to find next free slot")
var errJoinQueue = errors.New("Failed to join the queue")

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
	exists, err := s.userQueueNumberRepo.existsByUsername(username)
	if err != nil {
		log.Println("Error checking if user have already joined a queue:", err)
		return errUserQueueCheck
	}
	if exists {
		return errUserJoinedQueue
	}

	q, err := s.repo.getById(queueId)
	if err != nil {
		log.Println("Error getting queue by id", err)
		return errNextFreeCheck
	}

	tx, err := db.Db().Begin()
	if err != nil {
		log.Println("Error starting a transaction:", err)
		return errJoinQueue
	}

	if err = s.repo.incrementNextFreeSlot(tx, queueId); err != nil {
		log.Println("Error incrementing queue's next free slot number:", err)
		tx.Rollback()
		return errJoinQueue
	}

	uqn := userQueueNumber{
		Username: username,
		QueueId:  queueId,
		Number:   q.NextFreeSlotNumber,
	}
	if err = s.userQueueNumberRepo.save(tx, &uqn); err != nil {
		log.Println("Error creating user queue number record:", err)
		tx.Rollback()
		return errJoinQueue
	}

	if err = tx.Commit(); err != nil {
		log.Println("Error commiting transaction:", err)
		tx.Rollback()
		return errJoinQueue
	}

	return nil
}

func (s *service) getInfo(username string) (*userQueueNumber, error) {
	return s.userQueueNumberRepo.getByUsername(username)
}

func (s *service) next(username string) error {
	return s.repo.incrementCurrentSlot(username)
}
