package datastore

import (
	"github.com/codingchem/pokedexcli/internal/api"
	"github.com/codingchem/pokedexcli/internal/pokecache"
)

type ILocationStore interface {
	Next() ([]api.LocationArea, error)
	Prev() ([]api.LocationArea, error)
}

type LocationStore struct {
	next   *string
	prev   *string
	oldmap *api.Map
	cache  *pokecache.Cache
}

func (l *LocationStore) Next() ([]api.LocationArea, error) {
	return l.oldmap.NextLocations()
}

func (l *LocationStore) Prev() ([]api.LocationArea, error) {
	return l.oldmap.PreviousLocations()
}

func NewLocationStore() *LocationStore {
	return &LocationStore{
		next:   nil,
		prev:   nil,
		oldmap: &api.Map{},
		cache:  pokecache.NewCache(10),
	}
}
