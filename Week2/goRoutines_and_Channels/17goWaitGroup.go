package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	websites := []string{
		"http://google.com",
		"http://fb.com",
		"http://github.com",
		"http://linkedin.com",
		"http://twitter.com",
	}

	for _, v := range websites {
		// getStatusCode(v)
		go getStatusCode(v)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic("Opps Something wrong")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, url)
	}
}
