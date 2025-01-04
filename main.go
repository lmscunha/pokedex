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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	registry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for true {
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
