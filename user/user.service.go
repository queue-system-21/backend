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
	user := User{
		Username: username,
		Password: password,
	}
	return s.repo.save(&user)
}

func (s *Service) ValidateCredentials(username, password string) bool {
	user := User{
		Username: username,
		Password: password,
	}
	exists, err := s.repo.exists(&user)
	if err != nil {
		return false
	}
	return exists
}

func (s *Service) SetRole(tx *sql.Tx, username, code string) error {
	user := User{
		Username: username,
		RoleCode: code,
	}
	return s.repo.updateRole(tx, &user)
}

func (s *Service) GetRoleByUsername(username string) (string, error) {
	return s.repo.getRoleByUsername(username)
}
