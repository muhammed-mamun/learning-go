package main

import "fmt"

func main() {
	//uninitialized slices is nill
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//initialized slices can be generate with make func
	// var nums = make([]type, size, capacity)
	// var nums = make([]int, 0)
	// fmt.Println(nums == nil)
	// fmt.Println(nums)

	//capacity -> maximum number of elements can fit
	// fmt.Println(cap(nums))

	// nums = append(nums, 1)
	// nums = append(nums, 9)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 55)
	// nums = append(nums, 1)

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 55)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// slices - copy(fromarray, toarray)

	//slice operator
	var nums = []int{1, 2, 3}

	fmt.Println(nums[2:3])
}
