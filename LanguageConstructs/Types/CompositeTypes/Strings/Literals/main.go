package main

import "fmt"

// type alias str as string
type str string

func main() {

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
	// a Rune is an alias for uint32
	var rn = 'h'
	fmt.Println(rn)

	// `` is raw string literal syntax
	//
	var raw = `
hello world,\n
how are you?\t
{"_id": 1231231312,"data": [1, 2, 3, 4, 5],
	"used: false,
	"other": null
}
---
kind: deployment
`
	fmt.Print(raw)

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
	fmt.Println(symbol, string(symbol))

	// len on a string returns the number of bytes
	fmt.Println(len(str))
}
