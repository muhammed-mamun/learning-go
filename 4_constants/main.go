package main

import "fmt"

const age = 30

// constants - declared value can not be changed
func main() {
	const name string = "golang"

	// name = "js" can not be changed value

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port)
}
