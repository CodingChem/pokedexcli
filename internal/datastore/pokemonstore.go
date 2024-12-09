package datastore

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/codingchem/pokedexcli/internal/api"
	"github.com/codingchem/pokedexcli/internal/pokecache"
)

type IPokeStore interface {
	Catch(pokemon string) (caught bool, err error)
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
	data, err := api.GetPokemon(pokemon)
	if err != nil {
		return false, err
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
