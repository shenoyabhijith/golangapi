package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/shenoyabhijith/bookstore-api/models"
	"gorm.io/gorm"
)

type BookController struct {
	DB    *gorm.DB
	Store *sessions.CookieStore
	Tpl   *template.Template
}

func (bc *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	session, _ := bc.Store.Get(r, "session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var books []models.Book
	bc.DB.Find(&books)
	bc.Tpl.ExecuteTemplate(w, "books.html", books)
}

// Add other methods as needed
