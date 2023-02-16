package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if v == nil {
		return
	}
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		Logger.Info("Encoding error: %v", err)
		return
	}
}
