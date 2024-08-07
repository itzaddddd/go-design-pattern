package behavior

import "fmt"

type State interface {
	handle(ctx Context)
}

type Context struct {
	state State
}

func NewContext(state State) *Context {
	ctx := &Context{}
	ctx.TransitionTo(state)
	return ctx

}

func (c *Context) TransitionTo(state State) {
	c.state = state
	c.state.handle(*c)
}

func (c *Context) Request() {
	c.state.handle(*c)
}

type NewState struct{}

func NewNewState() *NewState {
	return &NewState{}
}

func (s *NewState) handle(ctx Context) {
	fmt.Println("user account is new state, activating now")
}

type ActiveState struct{}

func NewActiveState() *ActiveState {
	return &ActiveState{}
}

func (s *ActiveState) handle(ctx Context) {
	fmt.Println("user account is active, close account")
}

type CloseState struct{}

func NewCloseState() *CloseState {
	return &CloseState{}
}

func (s *CloseState) handle(ctx Context) {
	fmt.Println("user account is now closed")
}

func RunState() {
	ctx := NewContext(NewNewState())
	ctx.TransitionTo(NewActiveState())
	ctx.TransitionTo(NewCloseState())
}
