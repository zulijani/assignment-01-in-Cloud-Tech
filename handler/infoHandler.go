package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetInfoRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}

func handleGetInfoRequest(w http.ResponseWriter, r *http.Request) {
	countryCode := r.URL.Path[len("/countryinfo/v1/info/"):]

	limitStr := r.URL.Query().Get("limit")
	limit := 0 // default value of limit, 0 means no limit
	if limitStr != ""{
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter: "+err.Error(), http.StatusBadRequest)
			return
		}
	}
	res, err := http.Get(fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s?fields=name,continents,population,languages,borders,flags,capital", countryCode))
	if err != nil {
		http.Error(w, "Error fetching country info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(countryCode)

	defer res.Body.Close()


	var countryInfo CountryInfo

	if err := json.NewDecoder(res.Body).Decode(&countryInfo); err != nil {
		http.Error(w, "Error decoding country info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// checkign if the response contains at least one country
	/*if len(countryInfo) == 0{
		http.Error(w, "No country found for the provided code: "+countryCode, http.StatusNotFound)
		return
	}
	*/

	simplifiedCountryInfo := SimplifiedCountryInfo{
        Name:       countryInfo.Name.Common,
        Continents: countryInfo.Continents,
        Population: countryInfo.Population,
        Languages:  countryInfo.Languages,
        Borders:    countryInfo.Borders,
        Flag:       countryInfo.Flags.PNG,
        Capital:    countryInfo.Capital,
    }
	log.Printf("Fetching cities for country: %s", simplifiedCountryInfo.Name)
	

	encodedCountryName := url.QueryEscape(simplifiedCountryInfo.Name)
	response, err := http.Get(fmt.Sprintf("http://129.241.150.113:3500/api/v0.1/countries/cities/q?country=%s", encodedCountryName ))
    if err != nil {
        http.Error(w, "Error fetching country info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

	var citiesData struct {
		Data []string `json:"data"`
	}
	if err := json.NewDecoder(response.Body).Decode(&citiesData); err != nil {
		http.Error(w, "Error decoding cities: "+err.Error(), http.StatusInternalServerError)
		return
	}


	// aplying limit if specified
	if limit > 0 && limit < len(citiesData.Data){
		citiesData.Data = citiesData.Data[:limit]
	}



	simplifiedCountryInfo.Cities = citiesData.Data

	w.Header().Add("content-type", "application/json")

	//encoding the response and sending it back
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(simplifiedCountryInfo); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
