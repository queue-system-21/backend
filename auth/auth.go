package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"queue/db"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/sign-in", signIn).Methods(http.MethodPost)
	s.HandleFunc("/sign-up", signUp).Methods(http.MethodPost)
}

type authDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var dto authDto
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		fmt.Println("Sign in error:", err)
		http.Error(w, "Invalid request body", 400)
		return
	}
	fmt.Println(dto.Username, dto.Password)
	fmt.Fprintln(w, "you are signed in!")
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var dto authDto
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Invalid request body", 400)
		return
	}

	query := "insert into \"user\" (username, password) values ($1, $2);"
	if _, err := db.Db().Exec(query, dto.Username, dto.Password); err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Failed to create new user", 500)
		return
	}
}
