package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

		}
	}
}


func updateMovieVersion1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, _ := range movies {
		if movies[index].ID == params["id"] {
		_ = json.NewDecoder(r.Body).Decode(&movies[index])
			movies[index].ID = params["id"]
			json.NewEncoder(w).Encode(movies[index])

		}
	}
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "435673", Title: "Movie one", Director: &Director{FirstName: "Somesh", LastName: "Kanitkar"}})
	movies = append(movies, Movie{ID: "2", Isbn: "423567", Title: "Movie two", Director: &Director{FirstName: "Ashish", LastName: "Pawar"}})
	movies = append(movies, Movie{ID: "3", Isbn: "462776", Title: "Movie Three", Director: &Director{FirstName: "Rajat", LastName: "Achalkar"}})
	// for i:=4;i<1000000;i=i+1{
	// 	movies = append(movies, Movie{ID: "4", Isbn: "462776", Title: "Movie Three", Director: &Director{FirstName: "Rajat", LastName: "Achalkar"}})
	// }

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/moviesNewMethod/{id}", updateMovieVersion1).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Server running on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
