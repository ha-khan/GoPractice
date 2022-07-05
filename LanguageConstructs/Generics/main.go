package main

import "fmt"

type Added interface {
	int32 | float32 | string
}

func Adder[T Added](a, b T) T {
	return a + b
}

// Generic type
type Collection[T Added] struct {
	a, b T
}

func main() {
	for _, pair := range []Collection[int32]{{1, 2}, {3, 4}} {
		fmt.Println(Adder(pair.a, pair.b))
	}

	fmt.Println(Adder("hello ", "world"))

	var c any
	c = 123
	fmt.Println(c.(int) + 23)
}
