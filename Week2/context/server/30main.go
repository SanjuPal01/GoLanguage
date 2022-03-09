package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Panic(http.ListenAndServe("localhost:3000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	// case <-time.After(500 * time.Millisecond):
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
