package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		// time.Sleep(time.Second)    // if you wanna see the context cancelled error then uncomment this line and comment the next line
		time.Sleep(3 * time.Second)
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Hello")
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}
}
