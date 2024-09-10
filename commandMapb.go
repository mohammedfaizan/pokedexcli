package main

import (
	"fmt"
)

func commandMapb(cfg *Config, area ...string) error {

	if cfg.Next == "https://pokeapi.co/api/v2/location-area" || cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locationResp, err := cfg.pokeapiClient.ListLocationAreas(&cfg.Previous)
	if err != nil {
		return err
	}

	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}
	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous
	return nil
}
