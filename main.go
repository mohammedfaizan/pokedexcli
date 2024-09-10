package main

import (
	"time"

	"github.com/mohammedfaizan/pokedexcli/internal/pokeapi"
)

var config Config

func main() {
	config = Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		Next:          "https://pokeapi.co/api/v2/location-area",
		Previous:      "",
	}
	startRepl(&config)
}
