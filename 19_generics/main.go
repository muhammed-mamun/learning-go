package main

import (
	"fmt"
)

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	vals := []bool{true, false, false}
	nums := []int{15, 4, 5, 88, 6, 11}
	printSlice(vals)
	printSlice(nums)
}
