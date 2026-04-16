package user

import (
	"database/sql"
)

type Service struct {
	repo *repo
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (s *Service) Create(username, password string) error {
	return s.repo.create(username, password)
}

func (s *Service) ValidateCredentials(username, password string) bool {
	exists, err := s.repo.exists(username, password)
	if err != nil {
		return false
	}
	return exists
}

func (s *Service) SetRole(tx *sql.Tx, username, code string) error {
	return s.repo.setRole(tx, username, code)
}

func (s *Service) GetRoleByUsername(username string) (string, error) {
	return s.repo.getRoleByUsername(username)
}

func (s *Service) GetRoleByUsernameAndPassword(username, password string) (string, error) {
	return s.repo.getRoleByUsernameAndPassword(username, password)
}
