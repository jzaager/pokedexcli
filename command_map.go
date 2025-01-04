package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
