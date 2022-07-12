package Helpers

import "net/http"

func SetJsonHeaders(w http.ResponseWriter) {
	response := w
	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
}
