package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *gorm.DB
var err error

func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	db.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Use port 5432 if port-forwarding, or 30007 if using NodePort
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Book{})

	// Seed data if table is empty
	var count int64
	db.Model(&Book{}).Count(&count)
	if count == 0 {
		books := []Book{
			{Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
			{Title: "Introducing Go", Author: "Caleb Doxsey"},
			{Title: "Go in Action", Author: "William Kennedy"},
		}
		db.Create(&books)
	}

	http.HandleFunc("/books", getBooks)
	log.Println("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
