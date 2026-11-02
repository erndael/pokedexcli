package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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
	var locations LocationAreaResponse

	locationURL := baseURL + "/location-area/"
	if url != nil {
		locationURL = *url
	}

	data, ok := c.cache.Get(locationURL)
	if ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("could not decode response: %w", err)
		}
		return locations, nil
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

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("could not decode response: %w", err)
	}
	c.cache.Add(locationURL, data)

	return locations, nil
}
