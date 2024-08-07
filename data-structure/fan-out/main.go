package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	for {
		time.Sleep(time.Second)

		n := rand.Intn(100)

		fmt.Printf("Channel -> %d\n", n)
		ch <- n
	}

}

func consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("Consumer %v <- %d\n", name, n)
	}
}

func fanOut(chA <-chan int, chB, chC chan<- int) {
	for n := range chA {
		if n > 50 {
			chB <- n
		} else {
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA)
	go consumer(chB, "B")
	go consumer(chC, "C")

	fanOut(chA, chB, chC)
}
