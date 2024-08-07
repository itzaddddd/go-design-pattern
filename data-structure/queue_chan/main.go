package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	items chan int
	done  chan struct{}
	wg    sync.WaitGroup
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		items: make(chan int, bufferSize),
		done:  make(chan struct{}),
	}
}

func (q *Queue) Enqueue(value int) {
	q.wg.Add(1)
	q.items <- value
}

func (q *Queue) Dequeue() (val int, ok bool) {
	select {
	case val := <-q.items:
		q.wg.Done()
		return val, true
	case <-q.done:
		fmt.Print("done queue\n")
		return 0, false
	}

}
func (q *Queue) Close() {
	q.wg.Wait()
	close(q.done)
	close(q.items)
}

func main() {
	q := NewQueue(16)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	go func() {
		q.Close()
	}()

	for {
		val, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("%v\n", val)

	}

}
