package main

var config Config

func main() {
	config = Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	startRepl(&config)
}
