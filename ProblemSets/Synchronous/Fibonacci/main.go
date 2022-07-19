package main

import "fmt"

func main() {
	fmt.Println(fibonacci(70))
}

type cache map[int]int

var c = make(cache)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var left int
	if l, ok := c[n-1]; ok {
		left = l
	} else {
		c[n-1] = fibonacci(n - 1)
		left = c[n-1]
	}

	var right int
	if r, ok := c[n-2]; ok {
		right = r
	} else {
		c[n-2] = fibonacci(n - 2)
		right = c[n-2]
	}

	return left + right
}
