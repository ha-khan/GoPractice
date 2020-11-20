package main

import (
	"fmt"
)

/*
Write a function to replace all spaces in a string with `%20`

assume ASCII

rune(SPACE) -> 32
*/

const urlspace = "%20"
const asciispace = int32(32)

func urlify(str string) string {

	strcopy := str

	for idx := 0; idx < len(strcopy); {
		if rune(strcopy[idx]) == asciispace {
			strcopy = strcopy[:idx] + urlspace + strcopy[idx+1:]
			idx += 3
		} else {
			idx++
		}
	}

	return strcopy
}

func main() {
	fmt.Println(urlify("str s t r i n g"))
	fmt.Println(urlify("Mr John Smith  "))

}
