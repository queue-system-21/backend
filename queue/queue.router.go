package queue

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/queue").Subrouter()
	s.HandleFunc("/", list).Methods(http.MethodGet)
}
