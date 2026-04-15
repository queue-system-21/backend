package queue

import (
	"log"
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
