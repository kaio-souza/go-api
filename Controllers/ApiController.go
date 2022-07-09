package Controllers

import (
	"encoding/json"
	"go-api/Entities"
	"go-api/Helpers"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := Entities.ApiResponse{
		Message: "Test",
		Status:  "Status Test",
	}

	jsonResponse, err := json.Marshal(response)
	Helpers.ValidateErr(err)
	w.Write(jsonResponse)
	return

}
