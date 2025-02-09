package handler

import (
	"fmt"
	"net/http"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "text/html")

	output := `Welcome to the COUNTRY INFO SERVICE!!! BOM BOM BOM (FIREWORKS EXPLODING)!!
				Available endpoints:
				- GET /countryinfo/v1/info/{code} - Get country information
				- GET /countryinfo/v1/population/{code} - Get population data
				- GET /countryinfo/v1/status/ - Get service status
				`

	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}