package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"queue/db"

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

func createUser(username, password string) error {
	query := "insert into \"user\" (username, password) values ($1, $2);"
	_, err := db.Db().Exec(query, username, password)
	return err
}

func validateCredentials(username, password string) bool {
	query := `select exists(select *
              from "user"
              where username = $1
                and password = $2);`
	row := db.Db().QueryRow(query, username, password)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func createJwt(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	key := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(key))
}
