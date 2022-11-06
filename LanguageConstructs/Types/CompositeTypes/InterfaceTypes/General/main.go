package main

import (
	"fmt"
)

// Interface definition cannot be bound
// to a variable, can only be used as a
// https://go.dev/ref/spec#General_interfaces
type Network interface {
	~[4]byte | Uniform
	fmt.Stringer
}

// Defines a "Union" of types that some Generic Function/Type
type Uniform interface {
	~func() error | Tuple[string]
}

type Tuple[T any] struct {
	A, B T
}

type Node[T Network] struct {
	Value T
}

type fn func() error

func (f fn) String() string {
	return "function"
}

func main() {
	var n = Node[fn]{
		Value: fn(func() error { return nil }),
	}

	fmt.Println(n.Value.String())

}
