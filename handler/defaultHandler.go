package handler

import (
	"fmt"
	"net/http"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "text/html")

	output := "Welcome to the COUNTRY INFO SERVICE!!! BOM BOM BOM (FIREWORKS EXPLODING)!!<br>" +
				"Available endpoints:<br>" +
				"- GET /countryinfo/v1/info/{code} - Get country information<br>"+
				"- GET /countryinfo/v1/population/{code} - Get population data<br>"+
				"- GET /countryinfo/v1/status/ - Get service status"
				

	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}