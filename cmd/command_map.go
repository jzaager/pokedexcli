package cmd

import (
	"errors"
	"fmt"

	"github.com/jzaager/pokedexcli/config"
)

func MapF(cfg *config.Config, args ...string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationURL)
	if err != nil {
		return err
	}

	cfg.NextLocationURL = locationsResp.Next
	cfg.PrevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}

func MapB(cfg *config.Config, args ...string) error {
	if cfg.PrevLocationURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationURL)
	if err != nil {
		return err
	}

	cfg.NextLocationURL = locationsResp.Next
	cfg.PrevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}
