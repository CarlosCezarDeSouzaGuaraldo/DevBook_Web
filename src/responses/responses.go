package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error is the error response from API
type ErrorAPI struct {
	Err string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// FixStatusCodeError is a function to show the error type for client
func FixStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
