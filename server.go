package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/moficodes/bookdata/api/datastore"
	"github.com/moficodes/bookdata/api/loader"
	"log"
	"net/http"
	"time"
)

var (
	books datastore.BookStore
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	books = &datastore.Books{}
	books.Initialize()
}

func searchByISBN(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["isbn"]
	if ok {
		data := books.SearchISBN(val)
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				log.Fatalln(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func createBook(w http.ResponseWriter, r *http.Request) {
	ok := books.CreateBook(&loader.BookData{
		BookID:        "007",
		Title:         "NEW BOOK",
		Authors:       "MOFI",
		AverageRating: 0,
		ISBN:          "123123",
		ISBN13:        "123123123",
		LanguageCode:  "en",
		NumPages:      0,
		Ratings:       0,
		Reviews:       0,
	})
	w.Header().Set("Content-Type", "application/json")
	if ok {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success": "created"}`))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not created"}`))
}

func searchByAuthor(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["author"]
	if ok {
		data := *books.SearchAuthor(val)
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func searchByBookName(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["bookName"]
	if ok {
		data := *books.SearchBook(val)
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteByISBN(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet).q
	api.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	api.HandleFunc("/book", createBook).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}


