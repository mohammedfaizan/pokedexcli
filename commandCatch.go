package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no location provided")
	}
	pokemonName := args[0]
	pokemonInfoResp, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("throwing a ball at %s...\n", pokemonInfoResp.Name)

	if catchProbability(pokemonInfoResp.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("you may now inspect it with the inspect command.")
		cfg.caughtPokemon[pokemonName] = pokemonInfoResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func catchProbability(BaseExperience int) bool {
	randNum := rand.Intn(500)

	return randNum > BaseExperience
}
