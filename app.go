package main

import "fmt"

type node struct {
	data string
	next *node
}

type myQueueLinkedList struct {
	size int
	head *node
}

func (s *myQueueLinkedList) enqueue(name string) {
	newNode := &node{
		data: name,
	}
	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
	s.size++
}
func (s *myQueueLinkedList) dequeue() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("List is empty")
	}
	var prev *node
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	var data string
	if prev != nil {
		data = prev.next.data
		prev.next = nil
	} else {
		data = s.head.data
		s.head = nil
	}
	s.size--
	return data, nil
}
func (s *myQueueLinkedList) iterateList() {
	for node := s.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func newLinkedList() *myQueueLinkedList {
	return &myQueueLinkedList{}
}
func main() {
	queueList := newLinkedList()
	queueList.enqueue("1")
	queueList.enqueue("3")
	queueList.enqueue("5")
	queueList.iterateList()

	r, _ := queueList.dequeue()
	fmt.Println("Dequeue ", r)
	r1, _ := queueList.dequeue()
	fmt.Println("Dequeue ", r1)
	queueList.iterateList()
}
