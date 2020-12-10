package main

import "fmt"

func rotateMatrix(matrix ...[]int) [][]int {

	return nil
}

func main() {

	matrix := [][]int{[]int{1, 2}, []int{3, 4}}

	matrix = rotateMatrix(matrix...)

	fmt.Println(matrix)

	str := "hello"

	str2 := str[5:]

	fmt.Println(str2)

}
