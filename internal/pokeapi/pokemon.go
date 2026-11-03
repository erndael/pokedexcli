package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var pokemon Pokemon
	url := baseURL + "/pokemon/" + name

	if err := c.get(url, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
