package main

import (
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

func (s *SimpleContract)ConfirmOrder(res, id string)  {
	s.or = nil
}

func (s *SimpleContract)ServeOder(op, message string) error {
	s.or.HandleEvent(op, message)
	return nil
}

