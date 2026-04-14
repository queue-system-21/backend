package auth

import (
	"fmt"
	"log"
	"net/http"
)

func signIn(w http.ResponseWriter, r *http.Request) {
	dto, err := parseDto(r)
	if err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Invalid request body", 400)
		return
	}
	fmt.Println(dto.Username, dto.Password)
	fmt.Fprintln(w, "you are signed in!")
}

func signUp(w http.ResponseWriter, r *http.Request) {
	dto, err := parseDto(r)
	if err != nil {
		log.Println("Sign up error:", err)
		http.Error(w, "Invalid request body", 400)
		return
	}

	if err = createUser(dto.Username, dto.Username); err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Failed to create new user", 500)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintln(w, "User created successfully!")
}
