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
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := request.URL.Query()
	id := params.Get("id")
	var movie Movie
	for _, item := range movies {
		if item.ID == id {
			movie = item
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
	json.NewEncoder(writer).Encode(nil)
}

func addMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == movie.ID {
			movies[index] = movie
			json.NewEncoder(writer).Encode(movies[index])
		}
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := request.URL.Query()
	id := params.Get("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(writer).Encode(true)
			return
		}
	}
	json.NewEncoder(writer).Encode(false)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movie", getMovie).Methods("GET")
	router.HandleFunc("/addMovie", addMovie).Methods("POST")
	router.HandleFunc("/updateMovie", updateMovie).Methods("PUT")
	router.HandleFunc("/deleteMovie", deleteMovie).Methods("DELETE")

	// movies = append(movies, Movie{
	// 	ID:    "1",
	// 	Isbn:  "asdfg",
	// 	Title: "Naruto",
	// 	Director: &Director{
	// 		Firstname: "Debasish",
	// 		Lastname:  "Padhi",
	// 	},
	// })

	fmt.Printf("Starting server at port 6000 \n")
	log.Fatal(http.ListenAndServe(":6000", router))
}
