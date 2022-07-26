package main

import "fmt"

// any is an alias for interface{}
type Person[T any] struct {
	Friends []T
}

func (p *Person[any]) String() string {
	return fmt.Sprintf("%+v", p.Friends)
}

type friends struct{}

// interface is a type constraint
type Things interface {
	Added | friends
}

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

	p := Person[Person[int]]{
		Friends: make([]Person[int], 0),
	}

	fmt.Println(p)

	fmt.Println(Adder("hello ", "world"))

	var c any
	c = 123
	fmt.Println(c.(int) + 23)
}
