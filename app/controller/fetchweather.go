package controller

import (
	"encoding/json"
	"net/http"
)


func FetchWeather(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]bool{"fetch-weather": true})
}