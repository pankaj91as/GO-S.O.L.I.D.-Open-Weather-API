package main

import (
	"encoding/json"
)

type JSONParser interface {
    ParseJSON(data []byte) ([]City, error)
}

type CityJSONParser struct{}

func (p *CityJSONParser) ParseJSON(data []byte) ([]City, error) {
    var cities []City
    if err := json.Unmarshal(data, &cities); err != nil {
        return nil, err
    }
    return cities, nil
}