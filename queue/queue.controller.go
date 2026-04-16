package queue

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"queue/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func list(w http.ResponseWriter, r *http.Request) {
	queues, err := getAll()
	if err != nil {
		log.Println("Error in queue.list:", err)
		utils.SendErrMsg(w, "Error getting all queues", 500)
		return
	}

	if queues == nil {
		queues = make([]queue, 0)
	}
	if err := json.NewEncoder(w).Encode(queues); err != nil {
		log.Println("Error in queue.list:", err)
		utils.SendErrMsg(w, "Error getting all queues", 500)
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	dto, err := parseCreateDto(r)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to parse the input", 400)
		return
	}

	exists, err := existsByUserName(dto.ResponsibleUserUsername)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to create a queue", 500)
		return
	}
	if exists {
		utils.SendErrMsg(w, "You cannot assign this user for this queue", 400)
		return
	}

	if err = create(dto.Name, dto.ResponsibleUserUsername); err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to create a queue", 500)
		return
	}

	utils.SendSuccessMsg(w, "successfully created a queue", 201)
}

func delete(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, err := strconv.Atoi(pathVars["id"])
	if err != nil {
		log.Println("Error deleting a queue:", err)
		utils.SendErrMsg(w, "Invalid queue id", 400)
		return
	}
	if err = deleteById(id); err != nil {
		if errors.As(err, &errNoQueueDeleted{}) {
			utils.SendSuccessMsg(w, "No queue was deleted", 200)
			return
		}
		log.Println("Error deleting a queue", err)
		utils.SendErrMsg(w, "Error deleting a queue", 500)
		return
	}

	utils.SendSuccessMsg(w, "Successfully deleted the queue", 200)
}
