package pokeapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func Test() {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return
	}
	defer res.Body.Close()

	var locations LocationAreaResponse
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		log.Fatal(err)
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}
