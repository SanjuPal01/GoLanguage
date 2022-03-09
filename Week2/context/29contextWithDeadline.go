package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Hello")
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}
}
