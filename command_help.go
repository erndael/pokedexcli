package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for name, command := range config.commandRegistry {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}
