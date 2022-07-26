package main

import (
	"fmt"
	"sort"
)

// type alias str as string
type celsius float64

// Implements the stringer interface defined in the fmt package
// https://pkg.go.dev/fmt#Stringer
// similar notion of objects implementing __repr__ or __str__ in python
func (c celsius) String() string {
	return fmt.Sprintf("%.2f C", c)
}

func main() {

	var c celsius = 1.2
	fmt.Println(celsius(c))

	// "" is string literal syntax
	// Go source files are always encoded in UTF-8 and Go text strings are interpreted
	// as UTF-8, can include Unicode code points in string literals.
	//
	// "" escape characters have meaning and are processed accordingly
	var literal = "h\tel\tlo\n"
	fmt.Println(literal)

	// unicode escape sequence character; can repr a unicode symbol
	// by using ascii characters
	literal = "\u2665"
	fmt.Println(literal)
	fmt.Println(literal == "♥")

	// '' is rune literal syntax; the character code would show up
	// a rune is an alias for uint32
	var rn = 'h'
	fmt.Println(rn)

	// `` is raw string literal syntax, escape characters are not
	// processed, multi-lines are allowed
	var raw = `
hello world,\n
how are you?\t
{"_id": 1231231312,"data": [1, 2, 3, 4, 5],
	"used: false,
	"other": null
}
---
kind: deployment
\n
`
	fmt.Println(raw)

	// strings are considered an alias for []byte
	// indexing a str at some random location
	// will return a byte
	//
	// if only using ascii symbols this would be fine
	// unicode would return some blob of bits that are
	// a subset of the unicode char, UTF-8 ... which is variable length
	var str = "hello"
	symbol := str[0]
	fmt.Println(symbol, string(symbol))

	// len on a string returns the number of bytes
	fmt.Println(len(str))

	// strings are immutable, but can typecast to
	// []byte and mutate the underlying bytes
	// which would allow you to sort
	//
	//
	// however the underlying string needs to be all ascii
	// code points
	var unsorted = "gfedcba"
	var strByte = []byte(unsorted)
	sort.Slice(strByte, func(i, j int) bool {
		return string(strByte[i]) < string(strByte[j])
	})

	fmt.Println(string(strByte))
	fmt.Println(string(strByte) == "abcdefg")
}
