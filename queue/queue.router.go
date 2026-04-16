package queue

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/queue").Subrouter()
	s.Handle("/", newGetAllHandler()).Methods(http.MethodGet)
	s.Handle("/", newCreateHandler()).Methods(http.MethodPost)
	s.Handle("/{id:[0-9]+}", newDeleteHandler()).Methods(http.MethodDelete)
}
