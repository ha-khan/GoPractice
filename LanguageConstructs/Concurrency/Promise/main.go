package main

// Other aliases Future, Promise, Delay, Deferred
type Promise struct {
	state  string
	buffer []func() error
}

func (p *Promise) Then() *Promise {

	return p
}

func main() {
	p := Promise{}

	p.Then().Then().Then().Then()
}
