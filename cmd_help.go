package main

import (
	"fmt"

	"github.com/lmscunha/pokedexcli/internal/pokecache"
)

func commandHelp(cfg *config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCmds() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
