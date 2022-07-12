package Controllers

import (
	"encoding/json"
	gt "github.com/bas24/googletranslatefree"
	"go-api/Constants"
	"go-api/Entities"
	"go-api/Helpers"
	"io/ioutil"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	chuckResponse := getJoke()
	//jsonResponse, err := json.Marshal(response)

	response := Entities.ApiResponse{
		Message: translateJoke(chuckResponse.Joke),
	}
	Helpers.SetJsonHeaders(w)
	jsonResponse := Helpers.JsonEncode(response)

	w.Write(jsonResponse)
	return
}

func getJoke() Entities.JokeResponse {
	var responseObject Entities.JokeResponse

	response, _ := http.Get(Constants.ChuckNorrisEndpoint)
	responseData, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseData, &responseObject)
	return responseObject
}

func translateJoke(joke string) string {
	result, _ := gt.Translate(joke, "en", "pt-BR")
	return result
}
