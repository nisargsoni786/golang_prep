package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	params := mux.Vars(r)

	for _, item := range movies {
		// fmt.Println("item.Director --> ", item.Director)
		// fmt.Println("item.Director.f & item.Director.l", item.Director.Firstname, item.Director.Lastname)
		if item.ID == params["id"] {
			movie = item
			break

		}
	}

	json.NewEncoder(w).Encode(movie)
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

	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	// movie.ID = strconv.Itoa(rand.Intn(100000000))
	movie.ID = "3"
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	var new_movie Movie
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if item.ID == params["id"] {
			fmt.Println("Matched !!!", movie.Director)
			c_item := &movies[index]
			if movie.Isbn != "" {
				c_item.Isbn = movie.Isbn
			}
			if movie.Title != "" {
				c_item.Title = movie.Title
			}
			if movie.Director != nil && movie.Director.Firstname != "" {
				c_item.Director.Firstname = movie.Director.Firstname
			}
			if movie.Director != nil && movie.Director.Lastname != "" {
				c_item.Director.Lastname = movie.Director.Lastname
			}
			new_movie = *c_item
			break
		}
	}

	json.NewEncoder(w).Encode(new_movie)
}

func main() {
	r := mux.NewRouter()
	movies = append(
		movies,
		Movie{
			ID: "1", Isbn: "123", Title: "Pathaan",
			Director: &Director{Firstname: "Nisarg", Lastname: "Soni"},
		})
	movies = append(
		movies,
		Movie{
			ID: "2", Isbn: "456", Title: "Brahmashtra",
			Director: &Director{Firstname: "Nisarg", Lastname: "Soni-2"},
		})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
