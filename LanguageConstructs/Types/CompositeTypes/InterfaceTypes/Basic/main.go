package main

import (
	"fmt"
	"io"
)

// Interfaces are "abstract" types that define a "set of types"
// that implement the specification set by the interface definition
//
// A "Basic" interface is just a method set that some set of types
// each implement
//
// "Basic" interfaces can be bound to a variable name within some scope
// and act as containers for values (of some type) that implements all
// the methods of that interface
//
// https://go.dev/ref/spec#Basic_interfaces
type File interface {
	io.ReadCloser
	io.Writer
	fmt.Stringer
	error
	Name() string
	any
}

func main() {

}
