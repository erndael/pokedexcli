package pokeapi

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	BaseStat int              `json:"base_stat"`
	Stat     NamedAPIResource `json:"stat"`
}

type Type struct {
	Type NamedAPIResource `json:"type"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var pokemon Pokemon
	url := baseURL + "/pokemon/" + name

	if err := c.get(url, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
