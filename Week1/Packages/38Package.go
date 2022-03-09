package main

import (
	"fmt"
	"net/http"

	"example.com/sanju/util"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Length of 'Sanju' is %d\n", util.GetLength("Sanju"))
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":4000", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hii Sanju Here</h1>"))
}
