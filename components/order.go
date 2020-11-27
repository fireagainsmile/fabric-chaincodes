package components

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST0123456789")

type BusinessOrder struct {
	ID string
	OrderDetail string
	CurrentState StateHandlerInterface
	Done chan struct{}
	Err error
}

func NewOrder(or string) *BusinessOrder {
	state := GenerateStates()
	o := &BusinessOrder{
		ID: "123456",
		OrderDetail: or,
		CurrentState: state,
		Done: make(chan struct{}),
	}
	return o
}

func (o *BusinessOrder)HandleEvent(op , event string) *BusinessOrder {
	if o.CurrentState == nil {
		o.Err = errors.New("no procedure is ongoing")
		return o
	}
	err :=o.CurrentState.StateHandler(op, event)
	if err != nil {
		o.Err = err
	}
	if o.CurrentState.IsFinished() {
		next := o.CurrentState.Next()
		if next == nil {
			fmt.Println("Waiting for confirmation")
		}
		o.CurrentState = next
	}
	return o
}

func (o *BusinessOrder)Close()  {
	o.CurrentState = nil
	close(o.Done)
}

func (o *BusinessOrder)IsFinished() bool {
	return o.CurrentState == nil
}

func (o *BusinessOrder)GetStatus() string {
	if o.CurrentState != nil {
		return o.CurrentState.Name()
	}
	return "waiting for confirmation"
}

// help functions
func generateOrderID() string {
	return fmt.Sprintf("order-%s", randn(10))
}

func randn(n int) string {
	res := make([]rune, n)
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for i := range res {
		r := rand.Intn(len(letters))
		res[i] = letters[r]
	}
	return string(res)
}
