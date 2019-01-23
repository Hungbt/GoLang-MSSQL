package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type test struct {
	ID string `json:"name"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("GET")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var temp test
	_ = json.NewDecoder(r.Body).Decode(&temp)
	json.NewEncoder(w).Encode(temp)
}

// Main function
func main() {
	r := mux.NewRouter()

	// Route handles & endpoints
	go r.HandleFunc("/books", getBooks).Methods("GET")
	go r.HandleFunc("/books", createBook).Methods("POST")
	go r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	go r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	// Start server
	log.Fatal(http.ListenAndServe(":8000", corsOpts.Handler(r)))
}
