package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
    "net/url"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetPopulationRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}


func handleGetPopulationRequest(w http.ResponseWriter, r *http.Request) {
    // Extract the country code from the URL path
    countryCode := r.URL.Path[len("/countryinfo/v1/population/"):]

    // Fetch country name from the REST Countries API
    response, err := http.Get(fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s?fields=name", countryCode))
    if err != nil {
        http.Error(w, "Error fetching country info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    var countryName CountryName
    if err := json.NewDecoder(response.Body).Decode(&countryName); err != nil {
        http.Error(w, "Error decoding country info: "+err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println(countryName.Name.Common)
    // Parse the optional limit parameter (e.g., "2010-2015")
    limitParam := r.URL.Query().Get("limit")
    var startYear, endYear int
    if limitParam != "" {
        years := strings.Split(limitParam, "-")
        if len(years) != 2 {
            http.Error(w, "Invalid limit parameter. Expected format: startYear-endYear", http.StatusBadRequest)
            return
        }
        var err error
        startYear, err = strconv.Atoi(years[0])
        if err != nil {
            http.Error(w, "Invalid start year in limit parameter", http.StatusBadRequest)
            return
        }
        endYear, err = strconv.Atoi(years[1])
        if err != nil {
            http.Error(w, "Invalid end year in limit parameter", http.StatusBadRequest)
            return
        }
    }
    encodedCountryName := url.QueryEscape(countryName.Name.Common)
    res, err := http.Get(fmt.Sprintf("http://129.241.150.113:3500/api/v0.1/countries/population/q?country=%s", encodedCountryName))
    if err != nil {
        http.Error(w, "Error fetching country info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    var apiResponse APIResponse
    if err := json.NewDecoder(res.Body).Decode(&apiResponse); err != nil {
        http.Error(w, "Error decoding country info: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Filter population data based on the limit parameter
    var filteredData []PopulationData
    totalPopulation := 0
    count := 0

    for _, entry := range apiResponse.Data.PopulationCounts {
        if limitParam == "" || (entry.Year >= startYear && entry.Year <= endYear) {
            filteredData = append(filteredData, entry)
            totalPopulation += entry.Value
            count++
        }
    }

    // Calculate the mean population
    meanPopulation := 0
    if count > 0 {
        meanPopulation = totalPopulation / count
    }

    // Prepare the response
    resp := PopulationResponse{
        Mean:   meanPopulation,
        Values: filteredData,
    }

    // Send the response
    w.Header().Set("content-type", "application/json")
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
        return
    }
}