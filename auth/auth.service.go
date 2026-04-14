package auth

import (
	"encoding/json"
	"net/http"
	"queue/db"
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
