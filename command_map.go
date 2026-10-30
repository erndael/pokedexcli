package main

import "github.com/erndael/pokedexcli/internal/pokeapi"

func commandMap(config *config) error {
	pokeapi.Test()
	return nil
}
