package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

func newCustomer(name string, phone string) *customer {
	customerObj := customer{
		name:  name,
		phone: phone,
	}
	return &customerObj
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanoseconds precision
	customer
}

func neworder(id string, amount float32, status string, customer *customer) *order {
	//initial setup goes here
	myOrder := order{
		id:       id,
		amount:   amount,
		status:   status,
		customer: *customer,
	}
	return &myOrder
}

// // receiver type funcion
// // use aesteric when you need to change the value

// func (o *order) changeOrderStatus(status string) {
// 	o.status = status
// }

// // you don't need to use aesteric when you do not changing the value
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	//inline struct - similar to js object
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{
	// 	"golang",
	// 	true,
	// }

	// fmt.Println(language)

	customerObj := newCustomer("Ali", "0185465321")
	// fmt.Println(customerObj)
	myOrder := neworder("1", 30.60, "received", customerObj)
	fmt.Println(myOrder)

}
