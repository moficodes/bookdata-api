package datastore

import (
	"github.com/moficodes/bookdata/api/loader"
	"log"
	"os"
	"strings"
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

func (b *Books) SearchAuthor(author string) *[]*loader.BookData {
	ret := Filter(b.Store, func(v *loader.BookData) bool {
		return strings.Contains(strings.ToLower(v.Authors), author)
	})

	return ret
}

func (b *Books) SearchBook(bookName string) *[]*loader.BookData {
	ret := Filter(b.Store, func(v *loader.BookData) bool {
		return strings.Contains(strings.ToLower(v.Title), bookName)
	})

	return ret
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
