package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *Config) error {

	if cfg.Next == "https://pokeapi.co/api/v2/location-area" || cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := http.Get(cfg.Previous)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Println("error statuscode is:", resp.StatusCode)
	}
	if err != nil {
		fmt.Println(err)
	}

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return err
	}

	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
