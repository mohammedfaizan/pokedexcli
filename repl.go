package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var action []string
		if len(words) > 1 {
			action = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, action...)
			if err != nil {
				fmt.Println(err)
			}

			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(c *Config, a ...string) error { return commandHelp(c, a...) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(c *Config, a ...string) error { return commandExit(c, a...) },
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the pokemon world",
			callback:    func(c *Config, a ...string) error { return commandMap(c, a...) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas in the pokemon world",
			callback:    func(c *Config, a ...string) error { return commandMapb(c, a...) },
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Displays a list of all the pokemon in a given area",
			callback:    func(c *Config, a ...string) error { return explore(c, a...) },
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
