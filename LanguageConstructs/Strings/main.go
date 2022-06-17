package main

import "fmt"

// type alias str as string
type str string

func main() {

	var str = "hello"

	for _, symbol := range str {
		fmt.Println(symbol, string(symbol), byte(symbol))
	}

	// strings are considered an alias for []byte
	// indexing a str at some random location
	// will return a byte
	//
	// if only using ascii symbols this would be fine
	// unicode would return some blob of bits that are
	// a subset of the unicode char, UTF-8 ... which is variable length
	symbol := str[0]
	fmt.Println(symbol)
	fmt.Println(string(symbol))
}
