package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	// result := sum(1, 9, 5, 6, 3, 5, 8, 4, 2, 5, 6, 1)
	// fmt.Println(result)

	nums := []int{4, 88, 5, 2, 9, 6, 3, 11}
	reuslt := sum(nums...)

	fmt.Println(reuslt)
}
