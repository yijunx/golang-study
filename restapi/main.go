package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Init books slice
var books []Book

func main() {
	// init router
	r := mux.NewRouter()
	r.HandleFunc("/api/yo", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "masfasdgnuun98", Title: "book1", Author: Author{FirstName: "john", LastName: "doe"}})
	books = append(books, Book{ID: "2", Isbn: "uin89034fnnksd", Title: "book2", Author: Author{FirstName: "john2", LastName: "doe2"}})

	// route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	return
}

func createBook(w http.ResponseWriter, r *http.Request) {
	return
}
