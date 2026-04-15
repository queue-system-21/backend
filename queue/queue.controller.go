package queue

import (
	"encoding/json"
	"log"
	"net/http"
)

func list(w http.ResponseWriter, r *http.Request) {
	queues, err := getAll()
	if err != nil {
		log.Println("Error in queue.list:", err)
		http.Error(w, "Error getting all queues", 500)
		return
	}

	if err := json.NewEncoder(w).Encode(queues); err != nil {
		log.Println("Error in queue.list:", err)
		http.Error(w, "Error getting all queues", 500)
		return
	}
}
