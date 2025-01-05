package main

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func main() {

}
