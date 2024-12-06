package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationUrl string = "https://pokeapi.co/api/v2/location/"

type ApiResponse struct {
	Previous *string        `json:"previous"`
	Next     string         `json:"next"`
	Results  []LocationArea `json:"results"`
	Count    int            `json:"count"`
}
type LocationArea struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Map struct {
	Next     string
	Previous string
}

func (m *Map) NextLocations() ([]LocationArea, error) {
	if m.Next == "" {
		m.Next = locationUrl
	}
	res, err := http.Get(m.Next)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid error code: %d", res.StatusCode)
	}

	dec := json.NewDecoder(res.Body)
	var resData ApiResponse
	dec.Decode(&resData)
	m.Next = resData.Next
	// BUG: I have no idea if this is the proper check?
	if resData.Previous == nil {
		m.Previous = ""
	} else {
		m.Previous = *(resData.Previous)
	}
	return resData.Results, nil
}
