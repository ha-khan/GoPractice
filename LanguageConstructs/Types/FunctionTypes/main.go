package main

import "fmt"

func main() {

	//
	// closures
	// first-class
	// anonymous
	// variadic
	//
	//
	var t func(...int)
	switch t {
	case nil:
		fmt.Println("Is NULL")
	}
	t = func(i ...int) {}
	t(1, 2, 3, 4, 5, 6, 7, 8)

}
