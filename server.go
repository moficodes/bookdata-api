package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main(){
	log.Println("bookdata api")
	http.HandleFunc("/",home)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

