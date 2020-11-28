package main

import (
	"fmt"
)

/*

A string may encode an integer, e.g., "123" encodes 123. In this problem,
you are to implement methods that take a string representing an integer and
return the corresponding integer, and vice versa.

ASCII 0-9 [48-57]

*/
func convertStrToInt(integer string) int {

	var solution int

	for _, val := range integer {
		solution = solution*10 + (int(val) - 48)
	}

	return solution
}

func convertIntToString(integer int) string {
	var (
		solution string
		digits   = integer
	)

	for digits > 0 {
		solution = string(digits%10+48) + solution
		digits /= 10
	}

	return solution
}

func convertInputs(input ...interface{}) {
	for _, val := range input {
		switch concreteType := val.(type) {
		case string:
			fmt.Println(convertStrToInt(concreteType))
		case int:
			fmt.Println(convertIntToString(concreteType))
		default:
			panic(fmt.Sprintf("unknown type: %+v", concreteType))
		}
	}
}

func main() {

	input := make([]interface{}, 0)

	input = append(input, "123", 123, "342", 213)

	convertInputs(input...)

}
