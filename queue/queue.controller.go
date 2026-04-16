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

type getAllHandler struct {
	service *service
}

func newGetAllHandler() *getAllHandler {
	return &getAllHandler{
		service: newService(),
	}
}

func (g *getAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queues, err := g.service.getAll()
	if err != nil {
		log.Println("Error in queue.list:", err)
		utils.SendErrMsg(w, "Error getting all queues", 500)
		return
	}

	if err := json.NewEncoder(w).Encode(queues); err != nil {
		log.Println("Error in queue.list:", err)
		utils.SendErrMsg(w, "Error getting all queues", 500)
		return
	}
}

type createHandler struct {
	service *service
}

func newCreateHandler() *createHandler {
	return &createHandler{
		service: newService(),
	}
}

func (c *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := c.parseRequest(r)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to parse the input", 400)
		return
	}

	exists, err := c.service.existsByUsername(dto.ResponsibleUserUsername)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to create a queue", 500)
		return
	}
	if exists {
		utils.SendErrMsg(w, "You cannot assign this user for this queue", 400)
		return
	}

	if err = c.service.create(dto.Name, dto.ResponsibleUserUsername); err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to create a queue", 500)
		return
	}

	utils.SendSuccessMsg(w, "successfully created a queue", 201)
}

type createDto struct {
	Name                    string `json:"name"`
	ResponsibleUserUsername string `json:"responsibleUserUsername"`
}

func (c *createHandler) parseRequest(r *http.Request) (createDto, error) {
	var dto createDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	return dto, err
}

type deleteHandler struct {
	service *service
}

func newDeleteHandler() *deleteHandler {
	return &deleteHandler{
		service: newService(),
	}
}

func (d *deleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, err := strconv.Atoi(pathVars["id"])
	if err != nil {
		log.Println("Error deleting a queue:", err)
		utils.SendErrMsg(w, "Invalid queue id", 400)
		return
	}
	if err = d.service.deleteById(id); err != nil {
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
