package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

var startTime time.Time

func init(){
	startTime = time.Now()
}


func DiagnosticsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		handleGetDiagnosticsRequest(w)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}


func handleGetDiagnosticsRequest(w http.ResponseWriter){
	countriesNowStatus, err := checkAPIStatus("https://countriesnow.space/api/v0.1/countries/positions/q?iso2=NG")
	//doing this specific requests because it has small response

	if err != nil {
		http.Error(w, "Error checking CountriesNow API status: "+err.Error(), http.StatusInternalServerError)
		return
	}
	restCountriesStatus, err := checkAPIStatus("http://129.241.150.113:8080/v3.1/alpha/no?fields=continents")
	//doing this specific requests because it has small response


	if err != nil{
		http.Error(w, "Error checking RestCountries API status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	uptime := int64(time.Since(startTime).Seconds())

	res := DiagnosticsResponse {
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: restCountriesStatus,
		Version: 		  "v1",
		Uptime:			  uptime ,
	}


	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil{
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
        return
	}
}

func checkAPIStatus(url string) (int, error){
	res, err := http.Get(url)
	if err != nil{
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, nil
}