package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/moficodes/bookdata/api/datastore"
	"log"
	"net/http"
	"time"
)

var (
	books datastore.BookStore
	empty string = "{}"
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
	}
	w.WriteHeader(http.StatusNotFound)
}

func searchByBookName(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)
	val, ok := queries["book"]
	if ok {
		data := *books.SearchBook(val)
		b, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/authors/{author}", searchByAuthor).Methods("GET")
	api.HandleFunc("/books/{book}", searchByBookName).Methods("GET")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
