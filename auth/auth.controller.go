package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type signInResponse struct {
	Token string `json:"token"`
}

func signIn(w http.ResponseWriter, r *http.Request) {
	dto, err := parseDto(r)
	if err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Invalid request body", 400)
		return
	}

	userId, err := getUserId(dto.Username, dto.Password)
	if err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "User not found", 404)
		return
	}

	jwt, err := createJwt(userId)
	if err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Failed to sign in", 500)
		return
	}

	if err = json.NewEncoder(w).Encode(signInResponse{Token: jwt}); err != nil {
		log.Println("Sign in error:", err)
		http.Error(w, "Internal server error", 500)
		return
	}
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
