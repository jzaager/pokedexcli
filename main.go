package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	clean := strings.TrimSpace(text)
	lower := strings.ToLower(clean)
	words := strings.Split(lower, " ")
	return words
}
