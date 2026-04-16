package queue

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/queue").Subrouter()
	s.Handle("/", newGetAllHandler()).Methods(http.MethodGet)
	s.Handle("/", newCreateHandler()).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", delete).Methods(http.MethodDelete)
}
