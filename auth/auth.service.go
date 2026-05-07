package auth

import (
	"os"
	"queue/user"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	userService *user.Service
}

func newService() *service {
	return &service{
		userService: user.NewService(),
	}
}

func (s *service) createJwt(username, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(key))
}

func (s *service) validateCredentials(username, password string) bool {
	return s.userService.ValidateCredentials(username, password)
}

func (s *service) createUser(username, password string) error {
	return s.userService.Create(username, password)
}

func (s *service) getUserRole(username, password string) (string, error) {
	return s.userService.GetRoleByUsername(username)
}
