package main

import (
	// need to specify the module name ~ GoPractice
	// where GOROOT is auto-set
	"GoPractice/Tooling/Packages/internal/helpers"

	"fmt"
)

func main() {
	s, err := helpers.CopySlice([]int{1, 2, 3})
	if err != nil {
		fmt.Println(err.Error())
	}
	s = append(s, 4)
	fmt.Println(s)

	array := [3]int{1, 2, 3}
	_, err = helpers.CopySlice(array)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(array)
}
