package components

import (
	"fmt"
	"strings"
)

var DefaultOrg = NewOrg()

// restaurant registration and other supervise functions

type restaurantInfo struct {
	openning bool
	r *Restaurant
}

type Org struct {
	resInfo map[string] restaurantInfo
	order *OrderEvent
}

func NewOrg() *Org  {
	mapInfo := make(map[string]restaurantInfo)
	return &Org{
		resInfo: mapInfo,
	}
}
// functions used by restaurant
// register restaurant
func (o *Org)RegisterRestaurant(name string) *Org {
	id := randn(10)
	resID := fmt.Sprintf("shop-%s",id)
	res := NewRestaurant(name, resID)
	var resInfo restaurantInfo
	resInfo.openning = true
	resInfo.r = res
	o.resInfo[resID] = resInfo
	return o
}

func (o *Org)UpdateRestaurant(opening bool, id string) *Org {
	res, ok := o.resInfo[id]
	if !ok {
		fmt.Println("restaurant not exist")
		return o
	}
	res.openning = opening
	o.resInfo[id] = res
	return o
}

func (o *Org)ListRestaurant() string {
	var res []string
	for k, _ := range o.resInfo{
		res = append(res, k)
	}
	return strings.Join(res, ",")
}

func (o *Org)GetRestaurant(id string) *Restaurant {
	res, ok := o.resInfo[id]
	if !ok {
		fmt.Println("error: no matched restaurant found! ")
		return nil
	}
	return res.r
}


// restaurant operations
// commit an order
func (o *Org)DeliverOrder(want string, resId string) *OrderEvent {
	oe := NewOrderEvent(want)
	o.order = oe
	return  oe
}

// confirm order for users
func (o *Org)ConfirmOrder(resid string, orderId string)  {
	o.order = nil
}

func (o *Org)ServeOrder(op, message string)  {
	o.order.HandleEvent(op, message)
}

// cancel an order
// remove an order...



