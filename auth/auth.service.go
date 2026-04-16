package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type service struct{}

func newService() *service {
	return &service{}
}

func (s *service) createJwt(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	key := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(key))
}
