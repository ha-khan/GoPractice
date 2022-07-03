package main

import "fmt"

func main() {
	fmt.Println(fibonacci(40))
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

	return fibonacci(n-1) + fibonacci(n-2)
}
