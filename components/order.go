package components

import (
	"fmt"
	"math/rand"
	"time"
)


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRST0123456789")

type OrderEvent struct {
	ID string
	OrderDetail string
	CurrentState StateHandlerInterface
	Done chan struct{}
}

func NewOrderEvent(or string) *OrderEvent {
	state := GenerateStates()
	o := &OrderEvent{
		ID: generateOrderID(),
		OrderDetail: or,
		CurrentState: state,
		Done: make(chan struct{}),
	}
	return o
}

func (o *OrderEvent)HandleEvent(op , event string) *OrderEvent {
	o.CurrentState.StateHandler(op, event)
	if o.CurrentState.IsFinished() {
		next := o.CurrentState.Next()
		if next == nil {
			fmt.Println("Waiting for confirmation")
		}
		o.CurrentState = o.CurrentState.Next()
	}
	return o
}

func (o *OrderEvent)Close()  {
	o.CurrentState = nil
	close(o.Done)
}

func (o *OrderEvent)IsFinished() bool {
	return o.CurrentState == nil
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
