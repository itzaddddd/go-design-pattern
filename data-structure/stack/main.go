package main

import "fmt"

type Stack struct {
	itmes []int
}

func (s *Stack) Push(item int) {
	s.itmes = append(s.itmes, item)
}

func (s *Stack) Pop() int {
	lenStack := len(s.itmes)
	if lenStack == 0 {
		return -1
	}

	item, items := s.itmes[lenStack-1], s.itmes[0:lenStack-1]
	s.itmes = items

	return item
}

func main() {
	s := &Stack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)

	for {
		fmt.Printf("%v\n", s.Pop())
		if len(s.itmes) == 0 {
			break
		}

	}
}
