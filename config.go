package main

import "github.com/mohammedfaizan/pokedexcli/internal/pokeapi"

type Config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
}
