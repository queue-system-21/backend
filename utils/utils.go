package utils

import (
	"encoding/json"
	"net/http"
)

func SendSuccessMsg(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	jsonMsg := map[string]string{"message": msg}
	json.NewEncoder(w).Encode(jsonMsg)
}

func SendErrMsg(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	jsonMsg := map[string]string{"error": msg}
	json.NewEncoder(w).Encode(jsonMsg)
}
