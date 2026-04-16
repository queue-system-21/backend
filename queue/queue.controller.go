package queue

import (
	"encoding/json"
	"log"
	"net/http"
	"queue/utils"
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
