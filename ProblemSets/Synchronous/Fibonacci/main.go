package main

import (
	"fmt"
)

type implementer interface {
	Start() <-chan int
}

type FSM struct {
	orign func() bool
}

func main() {

	fmt.Println("Running")

	var f FSM

	f.orign = func() bool {
		return true
	}

	go f.orign()
}
