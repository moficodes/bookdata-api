package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/moficodes/bookdata/api/loader"
	"net/http"
	"strconv"
)

func searchByISBN(w http.ResponseWriter, r *http.Request) {
	queries := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	if val, ok := queries["isbn"]; ok {
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
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	ratingOver, ratingBelow, err := getRatingParams(r)
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	}
	if val, ok := pathParams["author"]; ok {
		data := *books.SearchAuthor(val, ratingOver, ratingBelow, limit, skip)
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
	w.WriteHeader(http.StatusNotFound)
}

func searchByBookName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ratingOver, ratingBelow, err := getRatingParams(r)
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	}
	pathParams := mux.Vars(r)
	if val, ok := pathParams["bookName"]; ok {
		data := *books.SearchBook(val, ratingOver, ratingBelow, limit, skip)
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
	w.WriteHeader(http.StatusNotFound)
}

func getRatingParams(r *http.Request) (float64, float64, error) {
	ratingOver := 0.0
	ratingBelow := 5.0
	queryParams := r.URL.Query()
	rOver := queryParams.Get("ratingOver")
	if rOver != "" {
		val, err := strconv.ParseFloat(rOver, 64)
		if err != nil {
			return ratingOver, ratingBelow, err
		}
		ratingOver = val
	}

	rBelow := queryParams.Get("ratingBelow")
	if rBelow != "" {
		val, err := strconv.ParseFloat(rBelow, 64)
		if err != nil {
			return ratingOver, ratingBelow, err
		}
		ratingBelow = val
	}
	return ratingOver, ratingBelow, nil
}

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}

func getSkipParam(r *http.Request) (int, error) {
	skip := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("skip")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return skip, err
		}
		skip = val
	}
	return skip, nil
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
