package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/shenoyabhijith/bookstore-api/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	store := sessions.NewCookieStore([]byte("super-secret-key"))
	tpl := template.Must(template.ParseGlob("views/*.html"))

	bookController := &controllers.BookController{
		DB:    db,
		Store: store,
		Tpl:   tpl,
	}

	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler(store, tpl))
	router.HandleFunc("/books", bookController.GetBooks).Methods("GET")

	log.Println("Server started at :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func loginHandler(store *sessions.CookieStore, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Same as before
	}
}
