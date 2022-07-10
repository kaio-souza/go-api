package Controllers

import (
	"encoding/json"
	"go-api/Entities"
	"go-api/Helpers"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := Entities.ApiResponse{
		Message: "Welcome to API",
		Secret:  "Secret Data",
	}

	jsonResponse, err := json.Marshal(response)
	Helpers.ValidateErr(err)
	w.Write(jsonResponse)
	return

}
