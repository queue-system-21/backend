package queue

import (
	"encoding/json"
	"log"
	"net/http"
	"queue/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type getAllHandler struct {
	service *service
}

func newGetAllHandler() http.Handler {
	return &getAllHandler{
		service: newService(),
	}
}

func (h *getAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queues, err := h.service.getAll()
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

func newCreateHandler() http.Handler {
	return &createHandler{
		service: newService(),
	}
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := h.parseRequest(r)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to parse the input", 400)
		return
	}

	role, err := h.service.getUserRole(dto.ResponsibleUserUsername)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to check assigned user role", 500)
		return
	}
	if role == "admin" {
		utils.SendErrMsg(w, "Cannot assign an admin", 400)
		return
	}

	err = h.service.existsBy(dto.ResponsibleUserUsername, dto.NameRus, dto.NameKaz)
	if err != nil {
		switch err {
		case errUserBusy:
			utils.SendErrMsg(w, "User is already assgined a queue", 400)
		case errNameRusNotUnique:
			utils.SendErrMsg(w, "nameRus is not unique", 400)
		case errNameKazNotUnique:
			utils.SendErrMsg(w, "nameKaz is not unique", 400)
		default:
			log.Println("Error creating a queue:", err)
			utils.SendErrMsg(w, "Failed to create a queue", 500)
		}
		return
	}

	err = h.service.create(dto.NameRus, dto.NameKaz, dto.ResponsibleUserUsername)
	if err != nil {
		log.Println("Error creating a queue:", err)
		utils.SendErrMsg(w, "Failed to create a queue", 500)
		return
	}

	utils.SendSuccessMsg(w, "successfully created a queue", 201)
}

type createDto struct {
	NameRus                 string `json:"nameRus"`
	NameKaz                 string `json:"nameKaz"`
	ResponsibleUserUsername string `json:"responsibleUserUsername"`
}

func (h *createHandler) parseRequest(r *http.Request) (createDto, error) {
	var dto createDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	return dto, err
}

type deleteHandler struct {
	service *service
}

func newDeleteHandler() http.Handler {
	return &deleteHandler{
		service: newService(),
	}
}

func (h *deleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, err := strconv.Atoi(pathVars["id"])
	if err != nil {
		log.Println("Error deleting a queue:", err)
		utils.SendErrMsg(w, "Invalid queue id", 400)
		return
	}
	if err = h.service.deleteById(id); err != nil {
		if err == errNoQueueDeleted {
			utils.SendSuccessMsg(w, "No queue was deleted", 200)
			return
		}
		log.Println("Error deleting a queue", err)
		utils.SendErrMsg(w, "Error deleting a queue", 500)
		return
	}

	utils.SendSuccessMsg(w, "Successfully deleted the queue", 200)
}
