package queue

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/queue").Subrouter()
	s.HandleFunc("/", list).Methods(http.MethodGet)
	s.HandleFunc("/", post).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", delete).Methods(http.MethodDelete)
}
