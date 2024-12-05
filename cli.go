package main

import (
	"os"
)

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func commandHelp() error {
	// TODO: Implement
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

var CommandMap = map[string]cliCommand{
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
