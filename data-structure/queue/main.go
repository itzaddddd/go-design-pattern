package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	lenQueue := len(q.items)
	if lenQueue == 0 {
		return -1
	}

	item, items := q.items[0], q.items[1:]
	q.items = items

	return item
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for {
		fmt.Printf("%v\n", q.Dequeue())
		fmt.Printf("left queue: %v\n", q)

		if len(q.items) == 0 {
			break
		}
	}

}
