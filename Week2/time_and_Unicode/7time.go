package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Println(currTime.Format("01-02-2006 Mon 15:04:05 MST")) // You need to do this way only

	createdDate := time.Date(2022, time.August, 01, 9, 12, 14, 34, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006"))

	fmt.Println(currTime.Weekday())
	fmt.Println(currTime.Add(time.Hour))
	fmt.Println(currTime.Clock())
	fmt.Println(currTime.UTC())
	fmt.Println(currTime.Equal(createdDate))
	fmt.Println(currTime.Location())
	fmt.Println(createdDate.Location())
}
