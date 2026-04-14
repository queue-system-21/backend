package auth

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/sign-in", signIn).Methods(http.MethodPost)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you are signed in!")
}
