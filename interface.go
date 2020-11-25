package main

import (
	"github.com/fireagainsmile/fabric-chaincodes/components"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleContract struct {
	contractapi.Contract
	org *components.Org
	users *components.User
}

// admin interfaces
func (s *SimpleContract)InitWithUser(name string) error {
	s.org = components.DefaultOrg
	s.users = components.NewUser(name)
	return nil
}

func (s *SimpleContract)RegisterRestaurant(name string) {
	s.org.RegisterRestaurant(name)
}

func (s *SimpleContract)UpdateRestaurant(openning bool, id string)  {
	s.org.UpdateRestaurant(openning, id)
}

func (s *SimpleContract)ListRestaurant()  string{
	return s.org.ListRestaurant()
}

func (s *SimpleContract)GetRestaurantName(id string) string  {
}

//user interfaces
func (s *SimpleContract)CommitOrder(want ,resId string) error {
	return s.users.CommitOrder(want, resId)
}

func (s *SimpleContract)ListOrders() string {
	return s.users.ListOrders()
}

func (s *SimpleContract)ConfirmOrder(res, id string)  {
	s.users.ConfirmOrder(res, id)
}

// restaurant interfaces
func (s *SimpleContract)QueryOderInfo() (string,error) {
	return  s.Name, nil
}

func (s *SimpleContract)ServeOder(op, message string) error {
	s.org.ServeOrder(op, message)
	return nil
}

