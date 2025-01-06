package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandMap(cfg *config) error {
	type result struct {
		Name string
	}

	type results []result

	type response struct {
		Next     string
		Previous string
		Results  results
	}

	res, err := http.Get(cfg.Next)
	if err != nil {
		fmt.Printf("error making get map request %v", err)
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
		fmt.Printf("error parsing get map response %v", err)
		os.Exit(1)
	}

	var bodyRes response
	if err = json.Unmarshal(body, &bodyRes); err != nil {
		fmt.Printf("error parsing get map body %v", err)
		os.Exit(1)
	}

	cfg.Next = bodyRes.Next
	cfg.Previous = bodyRes.Previous

	for _, location := range bodyRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
