package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "applicaton/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
