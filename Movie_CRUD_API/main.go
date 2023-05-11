package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"` //unique id with the movie
	Title    string    `json:"title"`
	Director *Director `josn:"director"`
}

type Director struct {
	FistName string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main() {
	fmt.Println("Welcome to movie API")
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
}
