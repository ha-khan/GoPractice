package main

import (
	"fmt"
	"strconv"
)

/*

Write a function that "compresses" string str : aabcccccaaa --> a2b1c5a3

if "compressed" string is greater than uncompressed, then return uncompressed string


*/

// Assume A-Z, a-Z
func compressString(str string) string {

	var (
		compressedStr string
		slow, fast    int
	)

	// Iterate slowly
	for slow = 0; slow < len(str) && fast < len(str); {
		var (
			currentChar rune
			counter     int
		)

		currentChar = rune(str[slow])

		// Iterate quickly
		for fast = slow; fast < len(str); fast++ {
			if currentChar != rune(str[fast]) {
				break
			}

			counter++
		}

		compressedStr += string(currentChar) + strconv.Itoa(counter)
		slow = fast
	}

	if len(compressedStr) < len(str) {
		return compressedStr
	}

}

func main() {

	str := compressString("aabcccccaaa")
	fmt.Println(str)

}
