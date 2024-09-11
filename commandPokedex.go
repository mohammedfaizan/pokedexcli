package main

import "fmt"

func pokedex(cfg *Config, args ...string) error {

	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}
