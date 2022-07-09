package Entities

type ApiResponse struct {
	Message string `json:"message"`
	Status  string `json:"-"`
}
