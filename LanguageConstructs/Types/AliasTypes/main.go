package main

import "fmt"

type Name string

type Age uint8

func (a Age) IsGreaterThan(other Age) bool {
	return a > other
}

type Attributes interface {
	~string | ~uint8
}

func Print[T Attributes](arg T) {
	fmt.Println(arg)
}

func main() {
	Print(Name("Name"))
	Print[uint8](12)
}
