package main

import (
	"fmt"
	"net/url"
)

func main() {
	newURl := url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	fmt.Println(newURl.String())
}
