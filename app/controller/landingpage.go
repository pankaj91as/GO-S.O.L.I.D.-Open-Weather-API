package controller

import (
	"encoding/json"
	"net/http"
)


func LandingPage(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]bool{"landing-page": true})
}