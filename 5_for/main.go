package main

import "fmt"

//in go 'for()' is the only contruct in go for looping
//there is no while(), no do while()
//but we can use for() loop as while

func main() {
	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// while() - infinite loop
	// for {
	// 	println('1')
	// }

	//classic for loop
	// for i := 0; i < 3; i++ {
	// 	// break

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//Range
	for i := range 11 {
		fmt.Println(i)
	}
}
