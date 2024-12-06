package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func commandHelp() error {
	fmt.Println("Welcome to the pokedex\n\nUsage:")
	for _, cmd := range CommandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func initCommands() {
	CommandMap = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

var CommandMap map[string]cliCommand
