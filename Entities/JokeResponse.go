package Entities

type JokeResponse struct {
	Id      string `json:"id"`
	IconUrl string `json:"icon_url"`
	Url     string `json:"url"`
	Joke    string `json:"value"`
}
