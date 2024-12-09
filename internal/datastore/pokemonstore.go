package datastore

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/codingchem/pokedexcli/internal/api"
	"github.com/codingchem/pokedexcli/internal/pokecache"
)

type IPokeStore interface {
	Catch(pokemon string) (caught bool, err error)
	Inspect(pokemon string) error
}

type PokemonStats struct {
	Stat  PokemonStat `json:"stat"`
	Value int         `json:"base_stat"`
}
type PokemonStat struct {
	Name string `json:"name"`
}

type PokemonTypesDetails struct {
	Types []PokemonTypes `json:"type"`
}
type PokemonTypes struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
	// TODO: fix stats and types
	// Types PokemonTypesDetails `json:"types"`
	// Stats      PokemonStats        `json:"stats"`
	Height     int `json:"height"`
	Weight     int `json:"weight"`
	Experience int `json:"base_experience"`
}
type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type pokemonStore struct {
	cache  *pokecache.Cache[[]byte]
	caught []string
}

func NewPokemonStore() *pokemonStore {
	store := pokemonStore{
		caught: make([]string, 1),
		cache:  pokecache.NewCache[[]byte](120),
	}
	return &store
}

func (ps *pokemonStore) Catch(pokemon string) (caught bool, err error) {
	data, success := ps.cache.Get(pokemon)
	if !success {
		data, err = api.GetPokemon(pokemon)
		if err != nil {
			return false, err
		}
		ps.cache.Add(pokemon, data)
	}
	var my_pokemon Pokemon
	err = json.Unmarshal(data, &my_pokemon)
	if err != nil {
		return false, err
	}
	// TODO: calculate based on base xp
	dc := 10
	dice := rand.New(rand.NewSource(time.Now().UnixNano()))
	if dice.Intn(20) > dc {
		ps.caught = append(ps.caught, pokemon)
		return true, nil
	}

	return false, nil
}

func (ps *pokemonStore) Inspect(pokemon string) (err error) {
	if !ps.isCaught(pokemon) {
		return fmt.Errorf("you have not caught that pokemon")
	}
	data, success := ps.cache.Get(pokemon)
	if !success {
		data, err = api.GetLocation(pokemon)
		if err != nil {
			return err
		}
		ps.cache.Add(pokemon, data)
	}
	var my_pokemon Pokemon
	err = json.Unmarshal(data, &my_pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", my_pokemon.Name, my_pokemon.Height, my_pokemon.Weight)
	// TODO: Print details of stats and types!
	return nil
}

func (ps *pokemonStore) isCaught(pokemon string) (caught bool) {
	for _, poke := range ps.caught {
		if pokemon == poke {
			return true
		}
	}
	return false
}
