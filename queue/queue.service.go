package queue

import (
	"encoding/json"
	"log"
	"net/http"
	"queue/db"
)

type queue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getAll() ([]queue, error) {
	query := "select id, name from queue"
	rows, err := db.Db().Query(query)
	if err != nil {
		return nil, err
	}

	var queues []queue
	for rows.Next() {
		var q queue
		if err = rows.Scan(&q.Id, &q.Name); err != nil {
			log.Println("Error in queue.getAll:", err)
			continue
		}
		queues = append(queues, q)
	}
	return queues, nil
}

func create(name, responsibleUserUsername string) error {
	query := `insert into queue (name, responsible_user_username) values ($1, $2)`
	_, err := db.Db().Exec(query, name, responsibleUserUsername)
	return err
}

func existsByUserName(username string) (bool, error) {
	query := `select exists(select id
              from queue
              where responsible_user_username = $1)`

	row := db.Db().QueryRow(query, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

type createDto struct {
	Name                    string `json:"name"`
	ResponsibleUserUsername string `json:"responsibleUserUsername"`
}

func parseCreateDto(r *http.Request) (createDto, error) {
	var dto createDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	return dto, err
}
