package main

import (
	"fmt"
	"github.com/fireagainsmile/fabric-chaincodes/components"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleContract struct {
	contractapi.Contract
	or *components.BusinessOrder
}

//user interfaces
func (s *SimpleContract)CommitOrder(want string) error {
	 s.or = components.NewOrder(want)
	return nil
}

func (s *SimpleContract)ListOrders() string {
	return s.or.ID
}

func (s *SimpleContract)GetStatus() string {
	if s.or != nil {
		return s.or.GetStatus()
	}else {
		return "Done"
	}
}

func (s *SimpleContract)ConfirmOrder(id string) string {
	if s.or == nil {
		return fmt.Sprintf("no order ongoing at this moment")
	}
	if s.or.ID != id {
		return fmt.Sprintf("no matched order found")
	}
	s.or = nil
	return "Confirmed"
}

func (s *SimpleContract)ServeOder(op, message string) error {
	s.or.HandleEvent(op, message)
	return nil
}

