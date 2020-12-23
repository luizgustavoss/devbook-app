package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseError represents response error from API
type ResponseError struct {
	Error string `json:"error"`
}

// JsonResponse returns a JSON response representation
func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil || statusCode == http.StatusNoContent{
		return
	}

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

// ErrorResponseResolver returns a json error representation
func ErrorResponseResolver(w http.ResponseWriter, r *http.Response) {
	var responseError ResponseError
	json.NewDecoder(r.Body).Decode(&responseError)
	JsonResponse(w, r.StatusCode, responseError)
}
