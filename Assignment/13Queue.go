package main

import (
	"fmt"

	"example.com/queue"
)

func main() {
	queue := queue.Init(1000)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println(queue.GetSize())
	fmt.Println(queue.GetFront())
	fmt.Println(queue.GetRear())
	queue.Dequeue()
	fmt.Println(queue.GetSize())
	fmt.Println(queue.GetFront())
	fmt.Println(queue.GetRear())
}
