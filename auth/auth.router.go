package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/sign-in", signIn).Methods(http.MethodPost)
	s.HandleFunc("/sign-up", signUp).Methods(http.MethodPost)
}
