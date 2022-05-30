package main

import "fmt"

type Compass string

const (
	NORTH Compass = "NORTH"
	SOUTH Compass = "SOUTH"
	WEST  Compass = "WEST"
	EAST  Compass = "EAST"
)

type Weekday int

// constant generat or iota , which is used to create a sequence
// of related values without spelling out each one explicitly.
// iota begins at zero and increments by one for each item in the sequence.
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	fn := func(w Weekday) {
		fmt.Println(w)
	}

	for _, val := range []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday} {
		fn(val)
	}

}
