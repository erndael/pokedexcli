package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	commandRegistry map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
				break
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

		if err := c.callback(config); err != nil {
			fmt.Println(err)
		}
	}
}
