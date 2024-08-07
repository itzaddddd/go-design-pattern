package behavior

import "fmt"

type Mediator interface {
	send(message string, colleague Colleague)
}

type Colleague interface {
	send(message string)
	receive(message string)
}

type AbstractColleague struct {
	mediator Mediator
}

func NewAbstractColleague(mediator Mediator) *AbstractColleague {
	return &AbstractColleague{mediator: mediator}
}

type ConcreteMediator struct {
	colleagues []Colleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

func (c *ConcreteMediator) register(colleague Colleague) {
	c.colleagues = append(c.colleagues, colleague)
}

func (c *ConcreteMediator) send(message string, colleague Colleague) {
	for _, col := range c.colleagues {
		if col != colleague {
			col.receive(message)
		}
	}
}

type ColleagueUser struct {
	AbstractColleague
	name string
}

func NewColleagueUser(mediator Mediator, name string) *ColleagueUser {
	return &ColleagueUser{
		AbstractColleague: AbstractColleague{
			mediator: mediator,
		},
		name: name,
	}
}

func (u *ColleagueUser) send(message string) {
	fmt.Printf("%s send message: %s", u.name, message)
	u.mediator.send(message, u)
}

func (u *ColleagueUser) receive(message string) {
	fmt.Printf("%s received message: %s", u.name, message)
}

func RunMediator() {
	mediator := NewConcreteMediator()

	john := NewColleagueUser(mediator, "John")
	jane := NewColleagueUser(mediator, "Jane")

	mediator.register(john)
	mediator.register(jane)

	john.send("Hi there!")
	jane.send("Hey!")
}
