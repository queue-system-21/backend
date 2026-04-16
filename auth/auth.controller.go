package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"queue/user"
	"queue/utils"
)

type signInResponse struct {
	Token string `json:"token"`
}

func signIn(w http.ResponseWriter, r *http.Request) {
	dto, err := parseDto(r)
	if err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Invalid request body", 400)
		return
	}

	if valid := user.ValidateCredentials(dto.Username, dto.Password); !valid {
		utils.SendErrMsg(w, "Credentials are invalid", 404)
		return
	}

	jwt, err := createJwt(dto.Username)
	if err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Failed to sign in", 500)
		return
	}

	if err = json.NewEncoder(w).Encode(signInResponse{Token: jwt}); err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Internal server error", 500)
		return
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	dto, err := parseDto(r)
	if err != nil {
		log.Println("Sign up error:", err)
		utils.SendErrMsg(w, "Invalid request body", 400)
		return
	}

	if err = user.Create(dto.Username, dto.Username); err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Failed to create new user", 500)
		return
	}

	utils.SendSuccessMsg(w, "User created successfully!", 201)
}
