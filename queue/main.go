package main

import "fmt"

type Queue struct {
	Capacity int
	Values   []int
}

func (q Queue) isEmpty() bool {
	return len(q.Values) == 0
}

func (q *Queue) Enquque(val int) {
	if len(q.Values) == q.Capacity {
		fmt.Println("Queue is full, cannot insert")
		return
	}

	q.Values = append(q.Values, val)
}

func (q *Queue) Dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}

	element := q.Values[0]
	if len(q.Values) == 1 {
		q.Values = nil
		return element
	}

	q.Values = q.Values[1:]
	return element
}

func (q Queue) Peek() int {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.Values[0]
}

func main() {
	q := Queue{Capacity: 5}
	fmt.Println(q.Values)

	for i := 0; i < 5; i++ {
		fmt.Printf("Enqueuing %d", i+1)
		q.Enquque(i + 1)
	}

	fmt.Println(q.Values)

	fmt.Printf("Dequeue: %d\n", q.Dequeue())
	fmt.Println(q.Values)
	fmt.Printf("Dequeue: %d\n", q.Dequeue())
	fmt.Println(q.Values)
	fmt.Printf("Dequeue: %d\n", q.Dequeue())
	fmt.Println(q.Values)

	fmt.Println(q.Peek())

	fmt.Println(q.Values)
}
