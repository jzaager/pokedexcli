package cmd

import (
	"errors"
	"fmt"

	"github.com/jzaager/pokedexcli/config"
)

func Inspect(cfg *config.Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("Must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("%q hasn't been caught yet!\n", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokeType.Type.Name)
	}

	fmt.Println()
	return nil
}
