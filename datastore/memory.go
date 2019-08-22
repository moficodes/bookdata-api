package datastore

import (
	"log"
	"os"
	"strings"

	"github.com/moficodes/bookdata/api/loader"
)

type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

func (b *Books) Initialize() {
	filename := "./assets/books.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}

func (b *Books) SearchAuthor(author string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.BookData {
	ret := Filter(b.Store, func(v *loader.BookData) bool {
		return strings.Contains(strings.ToLower(v.Authors), strings.ToLower(author)) && v.AverageRating > ratingOver && v.AverageRating < ratingBelow
	})
	if limit == 0 || limit > len(*ret) {
		limit = len(*ret)
	}
	data := (*ret)[skip:limit]
	return &data
}

func (b *Books) SearchBook(bookName string, ratingOver, ratingBelow float64, limit, skip int) *[]*loader.BookData {
	ret := Filter(b.Store, func(v *loader.BookData) bool {
		return strings.Contains(strings.ToLower(v.Title), strings.ToLower(bookName)) && v.AverageRating > ratingOver && v.AverageRating < ratingBelow
	})
	if limit == 0 || limit > len(*ret) {
		limit = len(*ret)
	}

	data := (*ret)[skip:limit]
	return &data
}

func (b *Books) SearchISBN(isbn string) *loader.BookData {
	ret := Filter(b.Store, func(v *loader.BookData) bool {
		return strings.ToLower(v.ISBN) == strings.ToLower(isbn)
	})
	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil
}

func (b *Books) CreateBook(book *loader.BookData) bool {
	*b.Store = append(*b.Store, book)
	return true
}

func (b *Books) DeleteBook(isbn string) bool {
	indexToDelete := -1
	for i, v := range *b.Store {
		if v.ISBN == isbn {
			indexToDelete = i
			break
		}
	}
	if indexToDelete >= 0 {
		(*b.Store)[indexToDelete], (*b.Store)[len(*b.Store)-1] = (*b.Store)[len(*b.Store)-1], (*b.Store)[indexToDelete]
		*b.Store = (*b.Store)[:len(*b.Store)-1]
		return true
	}
	return false
}

func (b *Books) UpdateBook(isbn string, book *loader.BookData) bool {
	for _, v := range *b.Store {
		if v.ISBN == isbn {
			v = book
			return true
		}
	}
	return false
}

func Filter(vs *[]*loader.BookData, f func(*loader.BookData) bool) *[]*loader.BookData {
	vsf := make([]*loader.BookData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}
