package datastore

import (
	"github.com/moficodes/bookdata/api/loader"
	"log"
	"os"
)

type books struct {
	store    *[]*loader.BookData
	filename string
}

func (b *books) Initialize() {
	file, err := os.Open(b.filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
}

func (b *books) SearchAuthor(author string) *[]*loader.BookData {
	return nil
}

func (b *books) SearchBook(bookName string) *[]*loader.BookData {
	return nil
}
