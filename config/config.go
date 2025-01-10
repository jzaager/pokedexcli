package config

import (
	"github.com/jzaager/pokedexcli/internal/pokeapi"
)

type Config struct {
	CaughtPokemon   map[string]pokeapi.Pokemon
	PokeapiClient   pokeapi.Client
	NextLocationURL *string
	PrevLocationURL *string
}
