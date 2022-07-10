package Entities

type ApiResponse struct {
	Message string `json:"message"`
	Secret  string `json:"-"`
}
