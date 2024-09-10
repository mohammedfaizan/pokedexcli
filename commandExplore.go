package main

import (
	"errors"
	"fmt"
)

func explore(cfg *Config, areas ...string) error {

	if len(areas) != 1 {
		return errors.New("no location provided")
	}
	locationAreaName := areas[0]
	locationResp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", locationResp.Name)
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
