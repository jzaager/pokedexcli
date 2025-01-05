package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide a location name")
	}

	locName := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetLocation(locName)
	if err != nil {
		return fmt.Errorf("Could not find any pokemon for area %s [name must match exactly as in the 'map' command, case insensitive]\n", locName)
	}

	fmt.Printf("\nExploring %s...\n", locName)
	fmt.Println("Found Pokemon:")
	for _, enc := range pokemonResp.PokemonEncounters {
		fmt.Println(" - ", enc.Pokemon.Name)
	}

	fmt.Println()
	return nil
}
