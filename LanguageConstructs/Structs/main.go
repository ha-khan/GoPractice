package main

import "fmt"

type Opts func()

type Interface interface {
	Config()
}

type Orchestrator struct {
	Name string

	// can embed a function type as a field in a struct, as they are considered "first class"
	// allows reassignment, dynamically when compared to a method
	config func(*Orchestrator, ...Opts)
	opts   []Opts
}

func (o *Orchestrator) Config() {
	o.config(o, o.opts...)
}

func main() {
	var o = new(Orchestrator)
	o.config = func(orch *Orchestrator, opts ...Opts) {
		fmt.Println("Done")
	}

	fn := func(i Interface) error {
		if _, ok := i.(*Orchestrator); ok {
			i.Config()
			return nil
		} else {
			return fmt.Errorf("Not *Orchestrator")
		}
	}

	fn(o)
}
