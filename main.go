package main

import (
	"fmt"
	"net/http"
	"queue/auth"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := mux.NewRouter()
	registerHealthHandler(r)
	auth.RegisterHandlers(r)
	http.ListenAndServe(":8080", r)
}

func registerHealthHandler(r *mux.Router) {
	s := r.PathPrefix("/health").Subrouter()
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The service is healthy")
	}
	s.HandleFunc("/", healthHandler).Methods(http.MethodGet)
}
