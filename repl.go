package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandRegistry() map[string]cliCommand {
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commandRegistry := getCommandRegistry()

	for {
		fmt.Print("Pokedex > ")

		if read := scanner.Scan(); !read {
			if err := scanner.Err(); err != nil {
				break
			}
			break
		}

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]
		c, ok := commandRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := c.callback(); err != nil {
			fmt.Println(err)
		}
	}
}
