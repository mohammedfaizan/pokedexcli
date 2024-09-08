package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *Config) error {

	resp, err := http.Get(cfg.Next)
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

	cfg.Next = locationResp.Next

	if prevUrl, ok := locationResp.Previous.(string); ok {
		cfg.Previous = prevUrl
	} else {
		cfg.Previous = ""
	}

	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}
	return nil
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
