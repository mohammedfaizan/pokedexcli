package main

import (
	"errors"
	"fmt"
)

func inspect(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no location provided")
	}
	pokemonName := args[0]

	if value, ok := cfg.caughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("Height: %d\n", value.Height)
		fmt.Printf("Weight: %d\n", value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
			switch stat.Stat.Name {
			case "hp":
				fmt.Printf("  -hp: %d\n", stat.BaseStat)
			case "attack":
				fmt.Printf("  -attack: %d\n", stat.BaseStat)
			case "defense":
				fmt.Printf("  -defense: %d\n", stat.BaseStat)
			case "special-attack":
				fmt.Printf("  -special-attack: %d\n", stat.BaseStat)
			case "special-defense":
				fmt.Printf("  -special-defense: %d\n", stat.BaseStat)
			case "speed":
				fmt.Printf("  -speed: %d\n", stat.BaseStat)
			}
		}
		fmt.Println("Types:")
		for _, typeInfo := range value.Types {
			fmt.Printf("   %s\n", typeInfo.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
