package queue

import (
	"net/http"
	"queue/middlewares"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/queue").Subrouter()
	s.Handle("", middlewares.NewRole([]string{"user", "admin"}, newGetAllHandler())).Methods(http.MethodGet)
	s.Handle("", middlewares.NewRole([]string{"admin"}, newCreateHandler())).Methods(http.MethodPost)
	s.Handle("/{id:[0-9]+}", middlewares.NewRole([]string{"admin"}, newDeleteHandler())).Methods(http.MethodDelete)
	s.Handle("/{id:[0-9]+}", newJoinHandler()).Methods(http.MethodPost)
	s.Handle("/number", newGetNumberHandler()).Methods(http.MethodGet)
	s.Handle("/next", middlewares.NewRole([]string{"receptionist"}, newNextHandler())).Methods(http.MethodPatch)

	s.Use(middlewares.NewAuthMiddleware)
}
