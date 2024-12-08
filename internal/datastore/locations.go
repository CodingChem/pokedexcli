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
	cache *pokecache.Cache[[]byte]
	next  string
}

func (l *LocationStore) Next() ([]LocationArea, error) {
	res, err := l.getData(l.next)
	if err != nil {
		return nil, err
	}
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
	res, err := l.getData(*l.prev)
	if err != nil {
		return nil, err
	}
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
		cache: pokecache.NewCache[[]byte](10),
	}
}

func (l *LocationStore) getData(url string) (api.ApiResponse, error) {
	data, err := api.GetLocations(url)
	if err != nil {
		return api.ApiResponse{}, err
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
