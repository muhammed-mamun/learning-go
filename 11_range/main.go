package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 6, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for i, num := range nums {

	// 	sum += num

	// 	fmt.Println(i, num)
	// }

	// fmt.Println(sum)

	// m := map[string]string{"fname": "john", "lname": "doe"}
	// key value
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// key
	// for k := range m {
	// 	fmt.Println(k)
	// }
	// value
	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
