package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	ID     string
	Title  string
	Year   int
	Genres []string
}

var movies []Movie

func main() {
	// read data from the CSV file
	file, err := os.Open("/home/kenny/Full stack/2/dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records[1:] {
		year, err := strconv.Atoi(record[4])
		if err != nil {
			log.Fatal(err)
		}
		genres := strings.Split(record[8], ",")
		movie := Movie{
			ID:     record[0],
			Title:  record[1],
			Year:   year,
			Genres: genres,
		}
		movies = append(movies, movie)
	}

	// set up HTTP handlers
	http.HandleFunc("/movies", handleMovies)
	http.HandleFunc("/movies/", handleMovie)
	http.HandleFunc("/search", handleSearch)

	// start the server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// return all movies
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%v", movies)
	case http.MethodPost:
		// create a new movie
		id := r.FormValue("_id")
		title := r.FormValue("title")
		year, err := strconv.Atoi(r.FormValue("year"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		genres := strings.Split(r.FormValue("genres"), ",")
		movie := Movie{
			ID:     id,
			Title:  title,
			Year:   year,
			Genres: genres,
		}
		movies = append(movies, movie)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleMovie(w http.ResponseWriter, r *http.Request) {
	// parse the movie ID from the URL
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	// find the movie with the given ID
	movie := findMovieByID(id)
	if movie == nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// return the movie
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%v", movie)
	case http.MethodPut:
		// update the movie
		if r.FormValue("title") != "" {
			movie.Title = r.FormValue("title")
		}
		year, err := strconv.Atoi(r.FormValue("year"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		movie.Year = year
		if r.FormValue("genres") != "" {
			genres := strings.Split(r.FormValue("genres"), ",")
			movie.Genres = genres
		}
		w.WriteHeader(http.StatusOK)
	case
