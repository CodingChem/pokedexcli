package datastore

import (
	"github.com/codingchem/pokedexcli/internal/pokecache"
)

type IPokeStore interface {
	Catch(pokemon string) (caught bool)
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
	Name       string              `json:"name"`
	Types      PokemonTypesDetails `json:"types"`
	Stats      PokemonStats        `json:"stats"`
	Height     int                 `json:"height"`
	Weight     int                 `json:"weight"`
	Experience int                 `json:"base_experience"`
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

func (ps *pokemonStore) Catch(pokemon string) (caught bool) {
	// TODO: false
	return false
}
