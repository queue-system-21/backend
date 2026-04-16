package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.Handle("/sign-in", newSignInHandler()).Methods(http.MethodPost)
	s.Handle("/sign-up", newSignUpHandler()).Methods(http.MethodPost)
}
