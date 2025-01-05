package main

import "fmt"

type orderStatus string

const (
	Received  orderStatus = "received"
	Confirmed orderStatus = "confirmed"
	Delivered orderStatus = "delivered"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changging order status to ", status)
}
func main() {
	changeOrderStatus(Delivered)
}
