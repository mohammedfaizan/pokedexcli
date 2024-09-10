package main

import (
	"fmt"
)

func commandMap(cfg *Config, areas ...string) error {

	locationResp, err := cfg.pokeapiClient.ListLocationAreas(&cfg.Next)
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
