package main

import (
	"fmt"
	"os"
	"os/exec"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()

	cmd := exec.Command("stty", "sane")
	cmd.Stdin = os.Stdin
	cmd.Run()

	os.Exit(0)
	return nil
}
