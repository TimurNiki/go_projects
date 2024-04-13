package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(status int, v any, w http.ResponseWriter) error{
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(status, map[string]string{"error": err.Error()}, w)
}

func ParseJSON(r *http.Request, v any) error{
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.body).Decode(v)
}