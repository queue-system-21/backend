package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type authDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func parseDto(r *http.Request) (authDto, error) {
	var dto authDto
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&dto)
	return dto, err
}

func createJwt(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	key := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(key))
}
