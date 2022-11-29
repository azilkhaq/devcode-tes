package helpers

import (
	"encoding/json"
	"net/http"
)

func Message(code string, message string) map[string]interface{} {
	return map[string]interface{}{"status": code, "message": message}
}

func Response(w http.ResponseWriter, code int, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}