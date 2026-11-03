package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/erndael/pokedexcli/internal/pokeapi"
)

type config struct {
	commandRegistry     map[string]cliCommand
	pokedex             map[string]pokeapi.Pokemon
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the 20 next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 previous locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location area>",
			description: "Explore an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "cath <pokemon>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if read := scanner.Scan(); !read {
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
			break
		}

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]
		c, ok := config.commandRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := c.callback(config, cleanedInput...); err != nil {
			fmt.Println(err)
		}
	}
}
