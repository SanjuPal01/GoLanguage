package main

import (
	"fmt"
	"time"
)

type Sleeper interface {
	Sleep()
}

type defaultSleeper struct {
}

func (d *defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Sum(x int, sleeper Sleeper) int {
	res := 0
	for i := 1; i <= x; i++ {
		// time.Sleep(1 * time.Second)
		sleeper.Sleep()
		res += i
	}
	return res
}

func main() {
	sleeper := &defaultSleeper{}
	fmt.Println(Sum(4, sleeper))
}
