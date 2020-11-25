package components

import (
	"errors"
	"strings"
)

type Restaurant struct {
	Name string
	ID string
	orders map[string]*OrderEvent
}



func NewRestaurant(name string, id string) *Restaurant {
	m := make(map[string]*OrderEvent)
	return &Restaurant{
		Name: name,
		ID: id,
		orders: m,
	}
	
}

// functions called by org
func (c *Restaurant)AddOrder(order *OrderEvent) error {
	if _, ok := c.orders[order.ID]; ok {
		return errors.New("order already exist! ")
	}
	c.orders[order.ID] = order
	return nil
}

func (c *Restaurant)QueryOrderInfo(id string) (string, error)  {
	order, ok := c.orders[id]
	if !ok {
		return "", errors.New("no matched order found! ")
	}
	return order.CurrentState.Name(), nil
}

func (c *Restaurant)ListAllOrders() string {
	var res []string
	for k, _ := range c.orders{
		res = append(res, k)
	}
	return strings.Join(res,",")
}

func (c *Restaurant)RemoveOrder(id string)  {
	delete(c.orders, id)
}

// functions used by restaurant employee
func (c *Restaurant)ServeOrder(orderId string, op string, message string) error {
	order, ok := c.orders[orderId]
	if !ok {
		return errors.New("Order does not exist! ")
	}
	if order.IsFinished() {
		return errors.New("Order is already finished! ")
	}
	if order.CurrentState == nil {
		return errors.New("[Error:] no handler for Order ")
	}

	order.HandleEvent(op, message)
	return nil
}