package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCmds() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
