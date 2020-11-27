package main

import (
	"fmt"
	"github.com/fireagainsmile/fabric-chaincodes/components"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleContract struct {
	contractapi.Contract
	or *components.BusinessOrder
	pm *components.ProcedureManager
}

//user interfaces
func (s *SimpleContract)Init()  {
	s.pm = components.DefaultManager
}

func (s *SimpleContract)CommitOrder(want string) error {
	 s.or = components.NewOrder(want)
	return nil
}

func (s *SimpleContract)ListOrders() string {
	if s.or == nil {
		return "no order exist"
	}
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
		fmt.Println("No order at the moment")
		return "no order waiting confirmation"
	}
	if s.or.ID == id{
		s.or = nil
		fmt.Println(id, ":Order Confirmed")
		return "Confirmed"
	}else {
		fmt.Println("no match order found", s.or.ID, id)
		return "order id does not match"
	}
}

func (s *SimpleContract)ServeOder(op, message string) error {
	s.or.HandleEvent(op, message)
	return nil
}

func (s *SimpleContract)CreateNewProcedure(procedureName, initialOp string) {
	s.pm.CreateNewProcedure(procedureName, initialOp)
}

func (s *SimpleContract)ConfigureMainOps(procedureName string, ops ...string)  {
	s.pm.ConfigureMainOps(procedureName, ops...)
}

func (s *SimpleContract)ConfigureSubs(procedureName, mainOpName string, subs ...string)  {
	s.pm.ConfigureSubs(procedureName, mainOpName, subs...)
}

func(s *SimpleContract)SetThreshHold(procedureName, opName string, threshHold int) {
	s.pm.SetThreshHold(procedureName, opName, threshHold)
}