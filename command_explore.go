package main

import "fmt"

func commandExplore(config *config, args ...string) error {
	if len(args) < 2 {
		fmt.Println("usage: explore <location name>")
		return nil
	}

	area, err := config.pokeapiClient.GetLocationArea(args[1])
	if err != nil {
		return err
	}
	for _, encounter := range area.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
