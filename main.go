package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
    {ID: "2", Title: "Introducing Go", Author: "Caleb Doxsey"},
    {ID: "3", Title: "Go in Action", Author: "William Kennedy"},
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func main() {
    http.HandleFunc("/books", getBooks)
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

