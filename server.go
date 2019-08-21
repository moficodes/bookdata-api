package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/moficodes/bookdata/api/datastore"
	"github.com/moficodes/bookdata/api/loader"
	"log"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")
	if ok {
		data := books.SearchISBN(val)
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}


func searchByAuthor(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["author"]
	w.Header().Set("Content-Type", "application/json")
	if ok {
		data := *books.SearchAuthor(val)
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func searchByBookName(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["bookName"]
	w.Header().Set("Content-Type", "application/json")
	if ok {
		data := *books.SearchBook(val)
		b, err := json.Marshal(data)
		if err != nil {
			w.Header()
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}


func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	avgRating, _ := strconv.ParseFloat(r.FormValue("AverageRating"), 64)
	numPages, _ := strconv.Atoi(r.FormValue("NumPages"))
	ratings, _ := strconv.Atoi(r.FormValue("Ratings"))
	reviews, _ := strconv.Atoi(r.FormValue("Reviews"))

	ok := books.CreateBook(&loader.BookData{
		BookID:        r.FormValue("BookID"),
		Title:         r.FormValue("Title"),
		Authors:       r.FormValue("Authors"),
		AverageRating: avgRating,
		ISBN:          r.FormValue("ISBN"),
		ISBN13:        r.FormValue("ISBN13"),
		LanguageCode:  r.FormValue("LanguageCode"),
		NumPages:      numPages,
		Ratings:       ratings,
		Reviews:       reviews,
	})
	if ok {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success": "created"}`))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not created"}`))
}

func deleteByISBN(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["isbn"]
	w.Header().Set("Content-Type", "application/json")
	if ok {
		deleted := books.DeleteBook(val)
		if deleted {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": "deleted"}`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "not deleted"}`))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet)
	api.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", searchByISBN).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	api.HandleFunc("/book", createBook).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}


