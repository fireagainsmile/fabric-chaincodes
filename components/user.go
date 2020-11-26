package components

import (
	"fmt"
)

type HosuseID int

type Property struct {
	MoneyAmount int
}

type User struct {
	Name string
	ID string
	orders map[string]*OrderEvent
	messageChan chan string
}

func NewUser(name string) *User {
	ch := make(chan string, 1)
	uID := fmt.Sprintf("user_%s", randn(10))
	return &User{
		Name :name,
		ID: uID,
		messageChan: ch,
		orders: make(map[string]*OrderEvent, 10),
	}
}

// ## Admin functions

// ## internal functions
// business logs here, ie. check auth, identity, user current data
func (u *User)rulesCheck()  {
	fmt.Println("business logic here")
}

func requireAuth()  {
	//
	fmt.Println("Checking auth here")
}
// go routine listening events when order is committed
//func (u *User) watch(oe *OrderEvent)  {
//	for  {
//		select {
//		case m := <- oe.Done:
//			fmt.Println("received message", m)
//			delete(u.orders, oe.ID)
//			return
//		default:
//			time.Sleep(10*time.Second)
//
//		}
//
//	}
//}

// functions used between modules
// communication between different modules
//func (u *User)sendOrder(message , resId string)  *OrderEvent{
//	// todo channel communication
//	return DefaultOrg.DeliverOrder(message, resId)
//}



//  interfaces for users module
//func (u *User)CommitOrder(want,  resId string) error {
//	u.rulesCheck()
//	//OE := u.SentOrder(want)
//	oe := u.sendOrder(want, resId)
//	if oe == nil {
//		return errors.New("failed to create order")
//	}
//	u.orders[oe.ID] = oe
//	return nil
//}
//
//func (u *User)ConfirmOrder(want string, orderId string)  {
//	DefaultOrg.ConfirmOrder(want, orderId)
//}
//
//func (u *User)ListOrders() string  {
//	var res []string
//	for _, v := range u.orders{
//		res = append(res, v.ID)
//	}
//	return strings.Join(res,",")
//}
//





