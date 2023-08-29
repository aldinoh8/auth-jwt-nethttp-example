package controller

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, responseBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseBody)
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}
