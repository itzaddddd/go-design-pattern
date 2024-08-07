package behavior

import (
	"fmt"
	"slices"
)

type Observer interface {
	update(subject Subject)
}

type Subject interface {
	attach(observer Observer)
	detach(observer Observer)
	notify()
}

type ConcreteSubject struct {
	observers []Observer
	state     int
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (c *ConcreteSubject) attach(observer Observer) {
	isExisted := slices.Contains(c.observers, observer)
	if isExisted {
		fmt.Println("this observer is already attached")
	}

	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) detach(observer Observer) {
	observerIndex := slices.Index(c.observers, observer)
	observers := slices.Delete(c.observers, observerIndex, observerIndex+1)
	c.observers = observers
}

func (c *ConcreteSubject) notify() {
	for _, observer := range c.observers {
		observer.update(c)
	}
}

func (c *ConcreteSubject) changeState(state int) {
	c.state = state
}

func (c *ConcreteSubject) getState() int {
	return c.state
}

type ConcreteObserver struct{}

func NewConcreteObserver() *ConcreteObserver {
	return &ConcreteObserver{}
}

func (c *ConcreteObserver) update(subject Subject) {
	concreteSubject, ok := subject.(*ConcreteSubject)
	if !ok {
		fmt.Println("subject failed")
	}

	state := concreteSubject.getState()
	if state < 3 {
		fmt.Println("concrete observer: reacted to the event")
	}

}
