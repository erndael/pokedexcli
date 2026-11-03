package main

import "fmt"

func commandHelp(config *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range config.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
