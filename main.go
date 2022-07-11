package main

import (
	"github.com/joho/godotenv"
	"go-api/Controllers"
	"go-api/Helpers"
	"net/http"
)

func main() {
	godotenv.Load()
	// create routes
	http.HandleFunc("/api", Controllers.Api)
	http.HandleFunc("/", Controllers.Home)

	// start server
	err := http.ListenAndServe(":3333", nil)
	Helpers.ValidateErr(err)
}
