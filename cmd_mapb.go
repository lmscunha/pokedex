package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lmscunha/pokedexcli/internal/pokecache"
)

func commandMapB(cfg *config, cache *pokecache.Cache) error {
	type result struct {
		Name string
	}

	type results []result

	type response struct {
		Next     string
		Previous string
		Results  results
	}

	url := cfg.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	urlData, ok := cache.Get(url)

	if ok == false {
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("error making get mapb request %v", err)
			os.Exit(1)
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			fmt.Printf("response fail with status code %d and\nbody: %s\n",
				res.StatusCode, body,
			)
			os.Exit(1)
		}

		if err != nil {
			fmt.Printf("error parsing get mapb response %v", err)
			os.Exit(1)
		}

		cache.Add(url, body)
		urlData, ok = cache.Get(url)
	}

	var mapData response
	if err := json.Unmarshal(urlData, &mapData); err != nil {
		fmt.Printf("error parsing get mapb body %v", err)
		os.Exit(1)
	}

	for _, location := range mapData.Results {
		fmt.Println(location.Name)
	}

	cfg.Previous = mapData.Previous
	cfg.Next = mapData.Next

	return nil
}
