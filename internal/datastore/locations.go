package datastore

import (
	"encoding/json"
	"fmt"

	"github.com/codingchem/pokedexcli/internal/api"
	"github.com/codingchem/pokedexcli/internal/pokecache"
)

type ILocationStore interface {
	Next() ([]LocationArea, error)
	Prev() ([]LocationArea, error)
	Get(area string) ([]Pokemon, error)
}

type LocationArea struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Pokemon struct {
	Name string `json:"name"`
}
type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LAApiResponse struct {
	Encounter []PokemonEncounter `json:"pokemon_encounters"`
}

type LocationStore struct {
	prev  *string
	cache *pokecache.Cache[[]byte]
	next  string
}

func (l *LocationStore) Next() ([]LocationArea, error) {
	res, err := l.getData(l.next)
	if err != nil {
		return nil, err
	}
	locs, err := l.unmarshal(res.Results)
	if err != nil {
		return nil, err
	}
	return locs, nil
}

func (l *LocationStore) Prev() ([]LocationArea, error) {
	if l.prev == nil {
		return nil, fmt.Errorf("Already on first page!")
	}
	res, err := l.getData(*l.prev)
	if err != nil {
		return nil, err
	}
	locs, err := l.unmarshal(res.Results)
	if err != nil {
		return nil, err
	}
	return locs, nil
}

func (l *LocationStore) Get(area string) ([]Pokemon, error) {
	data, err := api.GetLocation(area)
	if err != nil {
		return nil, err
	}
	var res LAApiResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	var pokemons []Pokemon
	for _, enc := range res.Encounter {
		pokemons = append(pokemons, enc.Pokemon)
	}
	return pokemons, nil
}

func NewLocationStore() *LocationStore {
	return &LocationStore{
		next:  "",
		prev:  nil,
		cache: pokecache.NewCache[[]byte](10),
	}
}

func (l *LocationStore) getData(url string) (api.ApiResponse, error) {
	data, ok := l.cache.Get(url)
	var err error
	if !ok {
		data, err = api.GetLocations(url)
		if err != nil {
			return api.ApiResponse{}, err
		}
		l.cache.Add(url, data)
	}
	var res api.ApiResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return api.ApiResponse{}, err
	}
	l.next = res.Next
	l.prev = res.Previous
	return res, nil
}

func (l *LocationStore) unmarshal(data []byte) ([]LocationArea, error) {
	var locs []LocationArea
	err := json.Unmarshal(data, &locs)
	if err != nil {
		return nil, err
	}
	return locs, nil
}
