package cli

import (
	"fmt"
	"os"

	"github.com/codingchem/pokedexcli/internal/datastore"
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

func mapCommand() error {
	curLocation, err := locations.Next()
	if err != nil {
		return err
	}
	for _, loc := range curLocation {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}

func mapbCommand() error {
	curLocation, err := locations.Prev()
	if err != nil {
		return err
	}
	for _, loc := range curLocation {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}

func initCommands() {
	locations = datastore.NewLocationStore()
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
		"map": {
			name:        "map",
			description: "Displays the location areas of the pokemon world, invoke again to paginate.",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Previous locations",
			callback:    mapbCommand,
		},
	}
}

var (
	CommandMap map[string]cliCommand
	locations  datastore.ILocationStore
)
