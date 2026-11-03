package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/erndael/pokedexcli/internal/pokeapi"
)

const catchThreshold = 60

func catchPokemon(config *config, name string) error {
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	caught := rand.IntN(pokemon.BaseExperience) < catchThreshold
	if !caught {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		addToPokedex(config, pokemon)
	}
	return nil
}

func addToPokedex(config *config, pokemon pokeapi.Pokemon) {
	config.pokedex[pokemon.Name] = pokemon
}

func commandCatch(config *config, args ...string) error {
	if len(args) < 2 {
		fmt.Println("usage: catch <pokemon name>")
		return nil
	}

	if err := catchPokemon(config, args[1]); err != nil {
		return err
	}

	return nil
}
