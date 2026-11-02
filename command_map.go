package main

import (
	"fmt"
)

func fetchAndPrintLocations(config *config, url *string) error {
	locations, err := config.pokeapiClient.GetLocationAreas(url)
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

func commandMap(config *config) error {
	return fetchAndPrintLocations(config, config.nextLocationURL)
}

func commandMapBack(config *config) error {
	if config.previousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchAndPrintLocations(config, config.previousLocationURL)
}
