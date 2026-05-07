package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"queue/auth"
	"queue/queue"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := mux.NewRouter()
	registerHealthHandler(r)
	auth.RegisterHandlers(r)
	queue.RegisterHandlers(r)
	port := os.Getenv("PORT")
	log.Println("Starting the server on port:", port)
	http.ListenAndServe(":"+port, r)
}

func registerHealthHandler(r *mux.Router) {
	s := r.PathPrefix("/health").Subrouter()
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The service is healthy")
	}
	s.HandleFunc("/", healthHandler).Methods(http.MethodGet)
}
