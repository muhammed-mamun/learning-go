package main

import "fmt"

// in go there are many forms of SWITCH statement

func main() {
	//simple switch

	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("One")

	// case 2:
	// 	fmt.Println("Two")

	// case 3:
	// 	fmt.Println("Three")

	// default:
	// 	fmt.Println("Other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's Weekend")
	// default:
	// 	fmt.Println("it's worked day")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		// switch i.(type) {}
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer - ", t)

		case string:
			fmt.Println("It's a string - ", t)

		case bool:
			fmt.Println("It's a boolean - ", t)
		default:
			fmt.Println("other - ", t)
		}
	}

	whoAmI("somaiya")
}
