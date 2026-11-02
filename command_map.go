package main

import (
	"fmt"
)

func fetchAndPrintLocations(config *config, url *string) error {
	locations, err := config.pokeapiClient.ListLocationAreas(url)
	if err != nil {
		return err
	}

	config.nextLocationURL = locations.Next
	config.previousLocationURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMap(config *config, args ...string) error {
	return fetchAndPrintLocations(config, config.nextLocationURL)
}

func commandMapBack(config *config, args ...string) error {
	if config.previousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchAndPrintLocations(config, config.previousLocationURL)
}
