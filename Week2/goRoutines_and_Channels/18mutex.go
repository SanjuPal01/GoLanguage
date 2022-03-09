package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}
var mut *sync.Mutex = &sync.Mutex{}
var statusCode = [][]string{}

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
	// fmt.Println(statusCode)
	wg.Wait()
	fmt.Println(statusCode)
}

func getStatusCode(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Opps Something Went Wrong")
	} else {
		mut.Lock()
		temp := []string{strconv.Itoa(res.StatusCode), url}
		statusCode = append(statusCode, temp)
		mut.Unlock()
	}
}
