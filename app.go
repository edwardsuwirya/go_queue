package main

import "fmt"

type queue []string

func (q *queue) enqueue(name string) {
	*q = append(*q, name)
}
func (q *queue) dequeue() string {
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}
func main() {
	q := make(queue, 0)
	q.enqueue("1")
	q.enqueue("3")
	q.enqueue("5")
	fmt.Println("List")
	fmt.Println(q)

	q.dequeue()
	q.dequeue()
	fmt.Println("List")
	fmt.Println(q)
}
