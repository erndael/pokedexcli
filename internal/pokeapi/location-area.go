package pokeapi

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetLocationAreas(url *string) (LocationAreaResponse, error) {
	locationURL := baseURL + "/location-area/"
	if url != nil {
		locationURL = *url
	}

	req, err := http.NewRequest(http.MethodGet, locationURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not get locations: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("%s: %s", locationURL, res.Status)
	}

	var locations LocationAreaResponse
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not decode response: %w", err)
	}

	return locations, nil
}
