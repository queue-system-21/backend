package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"queue/utils"
)

type requestParser struct{}

type authDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (rp *requestParser) parseDto(r *http.Request) (authDto, error) {
	var dto authDto
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&dto)
	return dto, err
}

type signInHandler struct {
	service *service
	*requestParser
}

func newSignInHandler() http.Handler {
	return &signInHandler{
		service: newService(),
	}
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *signInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := h.parseDto(r)
	if err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Invalid request body", 400)
		return
	}

	role, err := h.service.getUserRole(dto.Username, dto.Password)
	if err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Credentials are invalid", 404)
		return
	}

	jwt, err := h.service.createJwt(dto.Username, role)
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

type signUpHandler struct {
	service *service
	*requestParser
}

func newSignUpHandler() http.Handler {
	return &signUpHandler{
		service: newService(),
	}
}

func (h *signUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dto, err := h.parseDto(r)
	if err != nil {
		log.Println("Sign up error:", err)
		utils.SendErrMsg(w, "Invalid request body", 400)
		return
	}

	if err = h.service.createUser(dto.Username, dto.Username); err != nil {
		log.Println("Sign in error:", err)
		utils.SendErrMsg(w, "Failed to create new user", 500)
		return
	}

	utils.SendSuccessMsg(w, "User created successfully!", 201)
}
