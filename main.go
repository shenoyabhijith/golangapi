package main

import (
    "html/template"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    db    *gorm.DB
    err   error
    store = sessions.NewCookieStore([]byte("super-secret-key"))
    tpl   *template.Template
)

func init() {
    tpl = template.Must(template.ParseGlob("templates/*.html"))
}

type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Simple authentication logic
        if username == "admin" && password == "password" {
            session, _ := store.Get(r, "session")
            session.Values["authenticated"] = true
            session.Save(r, w)
            http.Redirect(w, r, "/books", http.StatusSeeOther)
            return
        }
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    tpl.ExecuteTemplate(w, "login.html", nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    var books []Book
    db.Find(&books)
    tpl.ExecuteTemplate(w, "books.html", books)
}

func main() {
    dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&Book{})

    router := mux.NewRouter()
    router.HandleFunc("/login", loginHandler)
    router.HandleFunc("/books", getBooks).Methods("GET")

    log.Println("Server started at :8081")
    log.Fatal(http.ListenAndServe(":8081", router))
}
 
