package main

import (
	"time"

	"github.com/erndael/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := config{
		commandRegistry:     getCommands(),
		pokeapiClient:       pokeClient,
		nextLocationURL:     nil,
		previousLocationURL: nil,
	}

	startRepl(&config)
}
