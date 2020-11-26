package main

import (
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
	return s.or.ID
}

func (s *SimpleContract)GetStatus() string {
	return s.or.CurrentState.Name()
}

func (s *SimpleContract)ConfirmOrder(id string)  {
	if s.or.ID == id{
		s.or = nil
		fmt.Println(id, ":Order Confirmed")
	}else {
		fmt.Println("no match order found")
	}
}

func (s *SimpleContract)ServeOder(op, message string) error {
	s.or.HandleEvent(op, message)
	return nil
}

