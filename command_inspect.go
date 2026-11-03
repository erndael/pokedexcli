package main

import (
	"fmt"

	"github.com/erndael/pokedexcli/internal/pokeapi"
)

func inspectStats(stats []pokeapi.Stat) {
	fmt.Println("Stats:")
	for _, stat := range stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
}

func inspectTypes(types []pokeapi.Type) {
	fmt.Println("Types:")
	for _, pkmType := range types {
		fmt.Printf("- %s\n", pkmType.Type.Name)
	}
}

func inspectPokemon(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d | Weight: %d\n", pokemon.Height, pokemon.Weight)
	inspectStats(pokemon.Stats)
	inspectTypes(pokemon.Types)
}

func commandInspect(config *config, args ...string) error {
	if len(args) < 2 {
		fmt.Println("usage: inspect <pokemon>")
		return nil
	}

	pokemon, ok := config.pokedex[args[1]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	inspectPokemon(pokemon)
	return nil
}
