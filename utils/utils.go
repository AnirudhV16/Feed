package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONParse(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing req body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteResponse(w http.ResponseWriter, status int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	WriteResponse(w, status, map[string]string{"error": err.Error()})
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
