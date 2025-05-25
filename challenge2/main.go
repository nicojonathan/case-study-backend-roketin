package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicojonathan/case-study-backend-roketin/challenge2/handler"
)

func main() {
	router := mux.NewRouter()

	// define endpoints for movies
	router.HandleFunc("/movies", handler.InsertMovie).Methods("POST")

	router.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")

	router.HandleFunc("/movies", handler.GetAllMovies).Methods("GET")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
}
