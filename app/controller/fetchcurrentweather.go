package controller

import (
	"encoding/json"
	"net/http"
)


func FetchCurrentWeather(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]bool{"fetch-current-weather": true})
}