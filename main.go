package main

import (
	"fmt"
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	slice := strings.Split(text, " ")

	slice = slices.DeleteFunc(slice, func(s string) bool {
		return s == ""
	})

	return slice
}

func main() {
	fmt.Println("Hello, World!")
}
