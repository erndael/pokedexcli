package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	if len(config.pokedex) < 1 {
		fmt.Println("you haven't caught any pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name, _ := range config.pokedex {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
