package main

import "github.com/mohammedfaizan/pokedexcli/internal/pokeapi"

type Config struct {
	caughtPokemon map[string]pokeapi.PokemonInfo
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
}
