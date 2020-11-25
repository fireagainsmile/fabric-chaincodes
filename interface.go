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
	return s.org.GetRestaurant(id).Name
}

//user interfaces
func (s *SimpleContract)CommitOrder(want string) error {
	return s.users.CommitOrder(want)
}

func (s *SimpleContract)ListOrders() string {
	return s.users.ListOrders()
}

func (s *SimpleContract)ConfirmOrder(want, id string)  {
	s.users.ConfirmOrder(want, id)
}

// restaurant interfaces
func (s *SimpleContract)QueryOderInfo(resId, orderId string) (string,error) {
	res := s.org.GetRestaurant(resId)
	return res.QueryOrderInfo(orderId)
}

func (s *SimpleContract)ServeOder(resId, orderId, op, message string) error {
	res := s.org.GetRestaurant(resId)
	return res.ServeOrder(orderId, op, message)
}

