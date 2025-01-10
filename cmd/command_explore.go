package cmd

import (
	"errors"
	"fmt"

	"github.com/jzaager/pokedexcli/config"
)

func Explore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide a location name")
	}

	locName := args[0]

	locationResp, err := cfg.PokeapiClient.GetLocation(locName)
	if err != nil {
		return fmt.Errorf("Could not find any pokemon for area %q [name must match exactly as in the 'map' command, case insensitive]\n", locName)
	}

	fmt.Printf("\nExploring %s...\n", locName)
	fmt.Println("Found Pokemon:")
	for _, enc := range locationResp.PokemonEncounters {
		fmt.Println(" - ", enc.Pokemon.Name)
	}

	fmt.Println()
	return nil
}
