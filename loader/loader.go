package loader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type BookData struct {
	BookID        string  `json:"book_id"`
	Title         string  `json:"title"`
	Authors       string  `json:"authors"`
	AverageRating float64 `json:"average_rating"`
	ISBN          string  `json:"isbn"`
	ISBN13        string  `json:"isbn_13"`
	LanguageCode  string  `json:"language_code"`
	NumPages      int     `json:"num_pages"`
	Ratings       int     `json:"ratings"`
	Reviews       int     `json:"reviews"`
}

func LoadData(fileName string) *[]*BookData {
	f, err := os.Open(fileName)
	reader := csv.NewReader(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	ret := make([]*BookData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		averageRating, _ := strconv.ParseFloat(row[3], 64)
		numPages, err := strconv.Atoi(row[7])
		if err != nil {
			log.Println(err)
		}
		ratings, err := strconv.Atoi(row[8])
		if err != nil {
			log.Println(err)
		}
		reviews, err := strconv.Atoi(row[9])
		if err != nil {
			log.Println(err)
		}
		book := &BookData{
			BookID:        row[0],
			Title:         row[1],
			Authors:       row[2],
			AverageRating: averageRating,
			ISBN:          row[4],
			ISBN13:        row[5],
			LanguageCode:  row[6],
			NumPages:      numPages,
			Ratings:       ratings,
			Reviews:       reviews,
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, book)
	}
	return &ret
}