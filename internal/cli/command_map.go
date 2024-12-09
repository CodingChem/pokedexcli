package cli

import (
	"fmt"
	"os"

	"github.com/codingchem/pokedexcli/internal/datastore"
)

type cliCommand struct {
	callback    func(string) error
	name        string
	description string
}

func commandHelp(_ string) error {
	fmt.Println("Welcome to the pokedex\n\nUsage:")
	for _, cmd := range CommandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(_ string) error {
	os.Exit(0)
	return nil
}

func mapCommand(_ string) error {
	curLocation, err := locations.Next()
	if err != nil {
		return err
	}
	for _, loc := range curLocation {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}

func mapbCommand(_ string) error {
	curLocation, err := locations.Prev()
	if err != nil {
		return err
	}
	for _, loc := range curLocation {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}

func exploreCommand(location string) error {
	pokemons, err := locations.Get(location)
	if err != nil {
		return err
	}
	for _, p := range pokemons {
		fmt.Printf("%v\n", p.Name)
	}
	return nil
}

func catchCommand(pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	caught, err := pokemons.Catch(pokemon)
	if err != nil {
		return err
	}
	switch caught {
	case true:
		fmt.Printf("%v was caught!\n", pokemon)
	case false:
		fmt.Printf("%v escaped!\n", pokemon)
	}
	return nil
}

func inspectCommand(pokemon string) error {
	err := pokemons.Inspect(pokemon)
	return err
}

func initCommands() {
	locations = datastore.NewLocationStore()
	pokemons = datastore.NewPokemonStore()
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
		"explore": {
			name:        "explore",
			description: "Explore the pokemon in the location area",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon!",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a pokemons details. Pokemon must first be caught",
			callback:    inspectCommand,
		},
	}
}

var (
	CommandMap map[string]cliCommand
	locations  datastore.ILocationStore
	pokemons   datastore.IPokeStore
)
