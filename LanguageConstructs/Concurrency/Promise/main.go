package main

import "fmt"

//
type Promise struct {
	state string // transition from ...

	// should make a (callback error) pair
	callbacks []func() error
}

func (p *Promise) Then(fn func() error) *Promise {

	return p
}

func (p *Promise) WhenError(fn func()) *Promise {
	return p
}

type Invoker struct {
	promise  *Promise
	executor func(*Promise)
}

func (i *Invoker) Register(fn func(*Promise)) *Promise {
	i.executor = fn
	i.promise = &Promise{
		state:     "INIT",
		callbacks: nil,
	}

	return i.promise
}

// will start the executor and monitor it for
func (i *Invoker) Start() {

}

func main() {
	p := Promise{}

	p.Then(func() error {
		fmt.Println("yo")
		return nil
	}).Then(
		func() error {
			return nil
		})
}
