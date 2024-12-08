package datastore

import (
	"encoding/json"

	"github.com/codingchem/pokedexcli/internal/api"
	"github.com/codingchem/pokedexcli/internal/pokecache"
)

type ILocationStore interface {
	Next() ([]LocationArea, error)
	Prev() ([]LocationArea, error)
}

type LocationArea struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type LocationStore struct {
	prev   *string
	oldmap *api.Map
	cache  *pokecache.Cache
	next   string
}

func (l *LocationStore) Next() ([]LocationArea, error) {
	res, err := l.oldmap.NextLocations()
	if err != nil {
		return nil, err
	}
	l.next = res.Next
	l.prev = res.Previous
	var locs []LocationArea
	err = json.Unmarshal(res.Results, &locs)
	if err != nil {
		return nil, err
	}
	return locs, nil
}

func (l *LocationStore) Prev() ([]LocationArea, error) {
	res, err := l.oldmap.PreviousLocations()
	if err != nil {
		return nil, err
	}
	l.next = res.Next
	l.prev = res.Previous
	var locs []LocationArea
	err = json.Unmarshal(res.Results, &locs)
	if err != nil {
		return nil, err
	}
	return locs, nil
}

func NewLocationStore() *LocationStore {
	return &LocationStore{
		next:   "",
		prev:   nil,
		oldmap: &api.Map{},
		cache:  pokecache.NewCache(10),
	}
}
