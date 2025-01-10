package main

import (
	"fmt"
	"strings"
)

func handleBackspace(buf *strings.Builder) {
	text := buf.String()
	buf.Reset()
	buf.WriteString(text[:len(text)-1])

	// Move cursor back, overwrite last char, and move back again
	fmt.Print("\b \b")
}

func handleEnter(buf *strings.Builder) (string, []string, error) {
	fmt.Println()
	text := buf.String()
	if text == "" {
		return "", nil, fmt.Errorf("No text provided")
	}

	cmd, args := splitCommandArgs(text)
	return cmd, args, nil
}

func splitCommandArgs(text string) (command string, args []string) {
	words := cleanInput(text)
	if len(words) == 0 {
		return "", nil
	}

	command = words[0]
	args = []string{}
	if len(words) > 1 {
		args = words[1:]
	}

	return command, args
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
