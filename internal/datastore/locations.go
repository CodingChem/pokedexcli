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
}

type LocationArea struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type LocationStore struct {
	prev  *string
	cache *pokecache.Cache[api.ApiResponse]
	next  string
}

func (l *LocationStore) Next() ([]LocationArea, error) {
	res, err := api.NextLocations(l.next)
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
	if l.prev == nil {
		return nil, fmt.Errorf("Already on first page!")
	}
	res, err := api.NextLocations(*l.prev)
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
		next:  "",
		prev:  nil,
		cache: pokecache.NewCache[api.ApiResponse](10),
	}
}
