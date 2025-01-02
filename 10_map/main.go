package main

import (
	"fmt"
	"maps"
)

func main() {
	// Map -> key value combination
	//creating map using make func
	// m := make(map[string]string)

	//setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// //get an element
	// fmt.Println(m["name"], m["area"])

	//IMP -> if key doesn't exist it return zeroed value
	// fmt.Println(m["book"])
	// fmt.Println(len(m))

	//remove element from map
	// delete(m, "area")
	// fmt.Println(m)

	// clear map
	// clear(m)

	// fmt.Println(m)

	m := map[string]int{"price": 350, "phones": 3}
	n := map[string]int{"price": 350, "phones": 3}

	fmt.Println(maps.Equal(m, n))

	v, ok := m["age"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

}
