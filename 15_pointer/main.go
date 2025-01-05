package main

import "fmt"

// by value - use value stored on memory
// func chnageNum(num int) {
// 	num = 5

// 	fmt.Println("In changeNum - ", num)
// }

func changeNum(num *int) {
	*num = 5
	println("in changeNum - ", *num)
}
func main() {
	num := 1

	changeNum(&num)

	fmt.Println("after change num - ", num)
}
