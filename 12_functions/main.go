package main

// func add(num1, num2 int) {
// 	sum := num1 + num2

// 	fmt.Println(sum)
// }

// //	func add(a int, b int) int {
// //		return a + b
// //	}
// func add(a, b int) int {
// 	return a + b
// }

// func getlanguages() (string, string, bool) {
// 	return "golang", "c", true
// }

func processIt(fn func(a int) int) {
	fn(8)
}

func main() {
	// result := add(4, 5)

	// fmt.Println(result)

	// fmt.Println(getlanguages())
	// lang1, lang2, lang3 := getlanguages()

	// fmt.Println(lang1, lang2, lang3)

	// lang1, lang2, _ := getlanguages()

	// fmt.Println(lang1, lang2)

	fn := func(a int) int {
		return 2
	}

	processIt(fn)
}
