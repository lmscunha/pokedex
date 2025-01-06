package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	slice := strings.Split(text, " ")

	return slices.DeleteFunc(slice, func(s string) bool {
		return s == ""
	})
}

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCmds() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous map locations",
			callback:    commandMapB,
		},
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getCmds()
	cfg := config{
		Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		parseIn := strings.Fields(strings.ToLower(input))
		cmd := registry[parseIn[0]]

		if value, exists := registry[cmd.name]; exists {
			if err := value.callback(&cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
