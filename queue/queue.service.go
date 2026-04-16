package queue

import (
	"encoding/json"
	"log"
	"net/http"
	"queue/db"
	"queue/user"
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
	tx, err := db.Db().Begin()
	if err != nil {
		return err
	}

	query := `insert into queue (name, responsible_user_username) values ($1, $2)`
	_, err = tx.Exec(query, name, responsibleUserUsername)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = user.SetRole(tx, responsibleUserUsername, "receptionist"); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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

type errNoQueueDeleted struct{}

func (n errNoQueueDeleted) Error() string {
	return "no queue was deleted"
}

func deleteById(id int) error {
	query := "delete from queue where id = $1"
	res, err := db.Db().Exec(query, id)
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return errNoQueueDeleted{}
	}
	return err
}
