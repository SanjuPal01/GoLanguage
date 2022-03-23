package queue

import "fmt"

type Queue struct {
	size int
	cap  int
	arr  []int
}

func Init(c int) *Queue {
	arr := make([]int, c)
	return &Queue{0, c, arr}
}

func (q *Queue) IsFull() bool {
	return q.size == q.cap
}
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Enqueue(x int) {
	if q.IsFull() {
		fmt.Println("Queue is Full")
		return
	}
	q.arr[q.size] = x
	q.size++
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
		return
	}
	for i := 0; i < q.size-1; i++ {
		q.arr[i] = q.arr[i+1]
	}
	q.size--
}

func (q *Queue) GetFront() int {
	if q.IsEmpty() {
		return -1
	} else {
		return q.arr[0]
	}
}

func (q *Queue) GetRear() int {
	if q.IsEmpty() {
		return -1
	} else {
		return q.arr[q.size-1]
	}
}

func (q *Queue) GetSize() int {
	return q.size
}
