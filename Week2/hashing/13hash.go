package main

import "fmt"

// var ArraySize int

type hashTable struct {
	// array [ArraySize]*bucket
	array []*bucket
}

type bucket struct {
	head *Node
}

type Node struct {
	key  string
	next *Node
}

func hash(key string, ArraySize int) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init(N int) *hashTable {
	// result := &hashTable{}
	result := &hashTable{
		array: make([]*bucket, N),
	}
	// (*result) is same as result
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (h *hashTable) Insert(key string, N int) {
	index := hash(key, N)
	h.array[index].insert(key)
}

func (h *hashTable) Search(key string, N int) bool {
	index := hash(key, N)
	return h.array[index].search(key)
}

func (h *hashTable) Delete(key string, N int) {
	index := hash(key, N)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		node := &Node{key: k}
		node.next = b.head
		b.head = node
	} else {
		fmt.Println("Key Already Exists")
	}
}

func (b *bucket) search(k string) bool {
	currNode := b.head
	for currNode != nil {
		if currNode.key == k {
			return true
		}
		currNode = currNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}

func main() {
	testHashTable := hashTable{}
	fmt.Println(testHashTable)

	var N int
	fmt.Scanln(&N)

	myHashTable := Init(N)
	// ArraySize = N
	fmt.Println(myHashTable)

	values := []string{
		"Sanju",
		"Sameer",
		"Nisha",
		"pooja",
	}

	for i := range values {
		myHashTable.Insert(values[i], N)
	}
	myHashTable.Insert("Nisha", N)

	fmt.Println(myHashTable.Search("Nisha", N))
	myHashTable.Delete("Nisha", N)
	fmt.Println(myHashTable.Search("Nisha", N))
}
