package pokeapi

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type LocationArea struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}

func (c *Client) ListLocationAreas(url *string) (LocationAreaList, error) {
	locationURL := baseURL + "/location-area/"
	if url != nil {
		locationURL = *url
	}

	var locations LocationAreaList
	if err := c.get(locationURL, &locations); err != nil {
		return LocationAreaList{}, err
	}

	return locations, nil
}

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	var location LocationArea
	if err := c.get(url, &location); err != nil {
		return LocationArea{}, err
	}

	return location, nil
}
