package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

Example 1:

Input: x = 123
Output: 321

*/
func main() {
	readArgsFromCLI()
	readFromStdin()
	readFromFile()
}

func reverse(x int) int {
	input := x
	var total int
	signed := 1
	if input < 0 {
		input *= -1
		signed = -1
	}

	for input > 0 {
		var temp int
		temp = input % 10
		total = total*10 + temp
		input /= 10
	}

	return total * signed

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func readArgsFromCLI() {
	var (
		input int
		err   error
	)

	// Read from CLI argument vector
	fmt.Printf("Input: %+v\n", os.Args)
	input, err = strconv.Atoi(os.Args[1])
	checkErr(err)
	total := reverse(input)
	fmt.Printf("Reverse: %+v\n", total)
}

func readFromStdin() {

	var (
		input int
		err   error
	)

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		input, err = strconv.Atoi(scanner.Text())
		checkErr(err)
		fmt.Printf("Input: %+v\n", input)
		fmt.Printf("Reverse: %+v\n", reverse(input))
	}

}

func readFromFile() {
	var (
		input  int
		stream *os.File
		err    error
	)
	stream, err = os.Open("data.txt")
	checkErr(err)
	for scanner := bufio.NewScanner(stream); scanner.Scan(); {
		input, err = strconv.Atoi(scanner.Text())
		checkErr(err)
		fmt.Printf("Input: %+v\n", input)
		fmt.Printf("Reverse: %+v\n", reverse(input))
	}

}
