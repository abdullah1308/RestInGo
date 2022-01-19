package main

import (
	// core package to work with json
	"encoding/json"

	// log to log errors
	"log"

	// to work with http apis
	"net/http"

	"math/rand"
	"strconv"

	// http router
	"github.com/gorilla/mux"
)

// Models
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
// Route handlers must have these two parameters passed to them
func getBooks(res http.ResponseWriter, req *http.Request) {
	// Setting the content type
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books)
}

// Get single book
func getBook(res http.ResponseWriter, req *http.Request) {
	// Setting the content type
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req) // Get params

	// Loop through books and find with id 
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

	json.NewEncoder(res).Encode(&Book{})
}

// Create a new book
func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(req.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID
	books = append(books, book)
	json.NewEncoder(res).Encode(book)
}


func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book

			_ = json.NewDecoder(req.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(res).Encode(book)
			return
		}
	}

	json.NewEncoder(res).Encode(books)
}

func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(books)
}

func main()  {
	// Init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
