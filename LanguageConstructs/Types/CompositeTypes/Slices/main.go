package main

import "fmt"

func main() {
	// initializes a slice of len = 0
	//
	// slices have three components: pointer, length, capacity
	var ints = make([]int, 0)

	for i := 0; i < 10; i++ {
		// append dynamically resizes the slice
		// by adding new elements
		ints = append(ints, i)
	}

	fmt.Println(ints)
	ints = append(ints, []int{10, 11, 12, 13, 14, 15}...)
	fmt.Println(ints)

	//
	fmt.Println(ints[0:2]) // prints idx 0, 1
	fmt.Println(ints[2:])
	fmt.Println(ints[:16])

	// the slice to copy values to needs to be the
	// exact length as the slice that elements are being
	// copied from
	//
	// otherwise zero values or not values would be added to
	// the copied slice
	var integers = make([]int, len(ints))
	copy(integers, ints)

	for idx, val := range integers {
		if val != ints[idx] {
			panic("error")
		}
	}
}
