package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"` //unique id with the movie
	Title    string    `json:"title"`
	Director *Director `josn:"director"`
}

type Director struct {
	Fistname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main() {
	fmt.Println("Welcome to movie API")

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "435", Title: "Fast And Furious", Director: &Director{Fistname: "Vin Diesel", Lastname: "Diesel"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1246", Title: "Need For Speed", Director: &Director{Fistname: "Thomas", Lastname: "Shelby"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at PORT :8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}
