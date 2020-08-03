package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book struct (model)
type Book struct {
	ID 		string 	`json:id`
	Isbn 	string 	`json:isbn`
	Title 	string 	`json:title`
	Author	*Author `json:author`
}

// Author struct

type Author struct {
	FirstName 	string `json:firstName`
	LastName	string `json:lastName`
}


var books []Book


// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update boook
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}




func main(){
	// Init router

	r := mux.NewRouter()

	// Mock Data 
	books = append(books, Book{ID: "1", Isbn: "44130", Title: "Book 1", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "73838", Title: "Book 2", Author: &Author{FirstName: "Mark", LastName: "Robin"}})

	// Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}