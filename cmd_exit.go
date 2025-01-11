package main

import (
	"fmt"
	"os"

	"github.com/lmscunha/pokedexcli/internal/pokecache"
)

func commandExit(cfg *config, cache *pokecache.Cache) error {
	cache.Stop()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
