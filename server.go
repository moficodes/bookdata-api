package main

import (
	"fmt"
	"github.com/moficodes/bookdata/api/loader"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	books *[]*loader.BookData
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init(){
	defer timeTrack(time.Now(), "file load")
	file, err := os.Open("assets/books.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	books = loader.LoadData(file)
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w,len(*books))
	fmt.Fprintln(w, "Hello")
}

func main(){
	log.Println("bookdata api")
	http.HandleFunc("/",home)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

