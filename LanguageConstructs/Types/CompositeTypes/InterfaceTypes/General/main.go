package main

import (
	"fmt"
)

// General interfaces combine both method definitions and some union of types
// and are meant to be used as "Type Constraints" for paramaterized arguments in Type/Function defintions
//
//	This interface definition cannot be bound to a variable, can only be used as a type constraint
//
// https://go.dev/ref/spec#General_interfaces
type Network interface {
	~[4]byte | Uniform // types in this "Network" set are meant to have either one as their "underlying structure"
	fmt.Stringer       // types in this "Network" set are meant to implement this method
}

// Defines a "Union" of types that some Generic Function/Type
type Uniform interface {
	~func() error | ~*Tuple[string] | ~chan Tuple[chan *Tuple[string]]
}

// Paramaterized Type definition
type Tuple[T any] struct {
	A, B T
}

func (t *Tuple[T]) String() string {
	return fmt.Sprintf("A: %v, B: %v", t.A, t.B)
}

type Node[T Network] struct {
	Value T
}

type fn func() error

func (f fn) String() string {
	return "function"
}

func Print[T Network](v T) {
	fmt.Println(v.String())
}

type receiver chan Tuple[chan<- *Tuple[string]]

func (r *receiver) String() string {
	return "uh"
}

func main() {
	var n = Node[fn]{
		Value: fn(func() error { return nil }),
	}

	var nn = Node[*Tuple[string]]{
		Value: &Tuple[string]{A: "Hello", B: "World"},
	}

	Print(n.Value)
	Print[*Tuple[string]](nn.Value)

	recursive := Tuple[*Tuple[string]]{
		A: &Tuple[string]{A: "Hello", B: "World"},
		B: &Tuple[string]{A: "How", B: "Are you?"},
	}

	Print(recursive.A)
	Print(recursive.B)
}
