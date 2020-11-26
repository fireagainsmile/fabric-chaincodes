package main

import (
	"errors"
	"fmt"
	"github.com/fireagainsmile/fabric-chaincodes/components"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleContract struct {
	contractapi.Contract
	or *components.OrderEvent
}

//user interfaces
func (s *SimpleContract)CommitOrder(want string) error {
	 s.or = components.NewOrderEvent(want)
	return nil
}

func (s *SimpleContract)ListOrders() string {
	if s.or != nil {
		return s.or.ID
	}
	return "nil"
}

func (s *SimpleContract)GetStatus() string {
	if s.or != nil {
		return s.or.CurrentState.Name()
	}else {
		return "Done"
	}
}

func (s *SimpleContract)ConfirmOrder(id string)  {
	if s.or == nil {
		fmt.Println("No order at the moment")
		return
	}
	if s.or.ID == id{
		s.or = nil
		fmt.Println(id, ":Order Confirmed")
	}else {
		fmt.Println("no match order found")
	}
}

func (s *SimpleContract)ServeOder(op, message string) error {
	if s.or == nil {
		return errors.New("No order at this moment ")
	}
	s.or.HandleEvent(op, message)
	return nil
}

