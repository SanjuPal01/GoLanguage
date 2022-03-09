package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second) // internally it call WithDeadline
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Hello")
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}
}
