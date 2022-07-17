package main

import "fmt"

func main() {
	// fixed sized
	//var compass = [4]string{"NORTH", "SOUTH", "WEST", "EAST"}
	var compass = [4]string{}

	var c = make(map[[4]string]int)

	c[compass]++

	fmt.Println(compass[0])
	fmt.Println(len(compass))
	fmt.Println(cap(compass))
}
