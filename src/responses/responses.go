package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func HandleStatusCodeErrors(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
