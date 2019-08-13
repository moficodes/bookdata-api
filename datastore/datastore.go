package datastore

import "github.com/moficodes/bookdata/api/loader"

type BookStore interface {
	Initialize()
	SearchAuthor(author string) *[]*loader.BookData
	SearchBook(bookName string) *[]*loader.BookData
}
