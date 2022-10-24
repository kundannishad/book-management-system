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
	Id       string    `json:"id,omitempty"`
	Isbn     string    `json:"isbn,omitempty"`
	Title    string    `json:"title,omitempty"`
	Director *Director `json:"director,omitempty"`
}

type Director struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	PhoneNo   string `json:"phone_no,omitempty"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("site-id", "asfsdt35346=dsfasd")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("site-id", "asfsdt35346=dsfasd")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("site-id", "asfsdt35346=dsfasd")
	params := mux.Vars(r)

	for _, Item := range movies {
		if Item.Id == params["id"] {
			json.NewEncoder(w).Encode(Item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("site-id", "asfsdt35346=dsfasd")

	fmt.Println("asfsdg", w)

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100022))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("site-id", "asfsdt35346=dsfasd")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

func main() {
	router := mux.NewRouter()

	//Movies Store Like this
	// movies = []Movie{{5, "2345456", "Movie Five", &Director{5, "Sandesh", "Singh", "sandeshsingh@gmail.com", "467677787878"}},
	// 	{6, "457457", "Movie Six", &Director{5, "Chand", "Singh", "chand@gmail.com", "57867978"}},
	// }

	// Apend Movie in movies
	movies = append(movies, Movie{"1", "12345678", "Movie One",
		&Director{"1", "Kundan", "kumar", "kundancse38@gmail.com", "8168149353"}})
	movies = append(movies, Movie{"2", "3453464", "Movie Two",
		&Director{"2", "Amit", "kumar", "kundancse38@gmail.com", "8168149353"}})
	movies = append(movies, Movie{"3", "12345678", "Movie Three",
		&Director{"3", "Akash", "kumar", "kundancse38@gmail.com", "8168149353"}})
	movies = append(movies, Movie{"4", "12345678", "Movie Four",
		&Director{"4", "Sachin", "kumar", "kundancse38@gmail.com", "8168149353"}})

	//Movies Store Like this
	// movies = []Movie{{5, "2345456", "Movie Five", &Director{5, "Sandesh", "Singh", "sandeshsingh@gmail.com", "467677787878"}},
	// 	{6, "457457", "Movie Six", &Director{5, "Chand", "Singh", "chand@gmail.com", "57867978"}},
	// }

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", getMoviesById).Methods("GET")
	router.HandleFunc("/movie/create", createMovie).Methods("POST")
	router.HandleFunc("/movie/update/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movie/delete/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server Start from port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
