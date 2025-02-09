package main

import(
	"a1/handler"
	"net/http"
	"log"
	"os"
)

func main(){

	port := os.Getenv("PORT")
	if port == ""{
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.INFO_PATH, handler.InfoHandler)
	

	log.Println("Starting server on port http://localhost:" + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
