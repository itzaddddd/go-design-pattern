package main

import "fmt"

type Iterator[T comparable] interface {
	next() T
	hasNext() bool
}

type Aggregate interface {
	createIterator() Iterator[any]
}

type Numbers struct {
	collection []int
}

func NewNumbers(collection []int) *Numbers {
	return &Numbers{
		collection: collection,
	}
}

func (n *Numbers) createIterator() Iterator[int] {
	return NewNumbersIterator(*n)
}

type NumbersIterator struct {
	collection Numbers
	position   int
}

func NewNumbersIterator(collection Numbers) *NumbersIterator {
	return &NumbersIterator{
		collection: collection,
	}
}

func (n *NumbersIterator) next() int {
	n.position++
	return n.collection.collection[n.position]
}

func (n *NumbersIterator) hasNext() bool {
	return n.position < len(n.collection.collection)
}

func RunIterator() {
	numbers := NewNumbers([]int{1, 2, 3, 4, 5})
	iterator := numbers.createIterator()

	for {
		if !iterator.hasNext() {
			break
		}

		fmt.Println(iterator.next())
	}
}

func main() {
	RunIterator()
}
