package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	citiesArray := getCitiesWithLatLong()

	// Print cities
	for _, city := range citiesArray {
		wg.Add(1)
		openWeatherAPI := &OpenWeatherAPI{}
		go openWeatherAPI.FetchData(city.Latitude, city.Longitude)
	}

	wg.Wait()
}

func getCitiesWithLatLong() (returnCity []City) {
	jsonParser := &CityJSONParser{}

	// Open JSON file
	jsonFile, err := os.Open("cities.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Read JSON File
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	// Parse JSON using the JSON parser
	cities, err := jsonParser.ParseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	return cities
}
