package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"queue/auth"
)

func main() {
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
	s.HandleFunc("/", healthHandler).Methods("GET")
}
