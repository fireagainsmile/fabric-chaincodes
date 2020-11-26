package components

import (
	"errors"
	"fmt"
	"time"
)



// handler interface for different states
type StateHandlerInterface interface {
	StateHandler(operation, message string) error
	Next() StateHandlerInterface
	Subs() []StateHandlerInterface
	IsFinished() bool
	Name() string
	Update()
}


type StateTemplate struct {
	op string
	log string
	done bool
	subStates []StateHandlerInterface
	ThreshHold int
	next StateHandlerInterface
	handler func(op, message string) error
}

func NewStateTemplate(op string, threshHold int) *StateTemplate {
	var subs []StateHandlerInterface
	return &StateTemplate{
		op: op,
		subStates: subs,
		ThreshHold: threshHold,
	}
}

// state template configuration
func (s *StateTemplate)AddSubs(handlerInterface ...StateHandlerInterface) *StateTemplate {
	s.subStates = append(s.subStates, handlerInterface...)
	return s
}

func (s *StateTemplate)SetNext(handlerInterface StateHandlerInterface) *StateTemplate {
	s.next = handlerInterface
	return s
}

func (s *StateTemplate)SetHandler(f func(string, string) error)  {
	s.handler = f
}


// interfaces of state handler
func (s *StateTemplate)Name() string  {
	return s.op
}

func (s *StateTemplate)IsFinished() bool  {
	return s.done
}

func (s *StateTemplate)Subs() []StateHandlerInterface  {
	return s.subStates
}

func (s *StateTemplate) Next() StateHandlerInterface  {
	return s.next
}

func (s *StateTemplate)StateHandler(op, message string) error  {
	EventCheck(message)
	var changed bool
	if s.done {
		fmt.Println("All done, nothing to do here")
		return  errors.New("nothing to do here")
	}
	if s.op == op {
		s.log = fmt.Sprintf("%s info: received message %s", time.Now().String(), message)
		if s.handler != nil {
			s.handler(op, message)
			changed = true
		}
		if changed{
			s.Update()
			return nil
		}
	}
	subs := s.Subs()
	if len(subs) != 0 {
		for _, ss := range subs{
			if ss.Name() == op {
				if ss.IsFinished() {
					fmt.Println("Sub operation is done!")
					return errors.New("Sub operation is already finished! ")
				}
				if err := ss.StateHandler(op, message); err != nil {
					return err
				}
				ss.Update()
				changed = true
			}
		}
	}
	if changed {
		s.Update()
		return nil
	}else {
		fmt.Println("no matched operation found!")
		return errors.New("no matched operation found")
	}

}


// help functions

func EventCheck(message string) error {
	return nil
}

func (s *StateTemplate)Update() {

	if s.ThreshHold < 0 {
		if checkAll(s) {
			s.done = true
		}
	}else if s.ThreshHold == 0 {
		s.done = true
	}else {
		if checkN(s, s.ThreshHold){
			s.done = true
		}
	}
}

func checkAll(handlerInterface StateHandlerInterface) bool {
	subs := handlerInterface.Subs()
	if len(subs) == 0 {
		return true
	}

	for _, v := range subs{
		if v.IsFinished() != true{
			return false
		}
	}
	return true
}

func checkN(handlerInterface StateHandlerInterface, n int) bool {
	subs := handlerInterface.Subs()
	if len(subs) < n {
		return false
	}
	counter := 0
	for _, v := range subs{
		if v.IsFinished() {
			counter ++
			if counter >= n{
				return true
			}
		}
	}
	return false
}


// business logic here
func GenerateStates() *StateTemplate {
	initState := NewStateTemplate("initialized", -1)
	waterState := NewStateTemplate("water", 0)
	flourState := NewStateTemplate("flour", 0)
	initState.AddSubs(waterState, flourState)
	f := func(op, message string)  error{
		// do nothing here
		return nil
	}
	initState.SetHandler(f)

	cookState := NewStateTemplate("waitingCooked", 1)
	woodState := NewStateTemplate("wood", 0)
	gasState := NewStateTemplate("gas", 0)
	cookState.AddSubs(woodState,gasState)

	deliverState := NewStateTemplate("deliver", 0)
	deliverState.SetNext(nil)

	// set up the state order here
	cookState.SetNext(deliverState)
	initState.SetNext(cookState)
	return initState

}