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

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getCmds()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		parseIn := strings.Fields(strings.ToLower(input))
		cmd := registry[parseIn[0]]

		if value, exists := registry[cmd.name]; exists {
			if err := value.callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
