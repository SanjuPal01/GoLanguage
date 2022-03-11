package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello All"))
	})

	http.HandleFunc("/hello", helloHandle)

	http.HandleFunc("/getbook", getBook)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Write([]byte("<h1>Hello</h1>"))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	book := Book{
		Title:  "The Girl on the Train",
		Author: "Sanju",
		Pages:  245,
	}

	json.NewEncoder(w).Encode(book)
}
