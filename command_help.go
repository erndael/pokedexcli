package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for name, command := range getCommandRegistry() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}
