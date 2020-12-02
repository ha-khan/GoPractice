package main

import (
	//"bytes"
	"fmt"
	//"os"
	"time"
)

// type keyword allows you to create your own user defined types
// similar to class in java

// struct, interface, string, map, int32, rune, byte .. are all built-in types in Go

type rawData []byte

type shifter func(data *byte)

func main() {

	dat := rawData{}

	dat = append(dat, byte(0))
	dat = append(dat, byte(1))
	dat = append(dat, byte(2))

	// Creates a variable of some type, and giving it the zero value for that type
	var shift shifter = func(data *byte) { *data = *data << 1 }

	for idx := range dat {

		fmt.Println(fmt.Sprintf("%08b", dat[idx]))

		shift(&dat[idx])
	}

	fmt.Println(fmt.Sprintf("%08b", dat))

	// unbuffered channel; empty struct doesn't take up space
	// <-chan string is a Read Only channel
	// chan<- string is a Write Only channel
	// chan string is a Read/Write channel
	gen := func() <-chan string {
		var c = make(chan string, 1)
		go func() { c <- "hello"; c <- "world"; c <- "!!!"; defer close(c) }()
		return c
	}

	// Will iterate forever doing the Read (<-) operation on the returned channel
	// from invoking gen() function
	for val := range gen() {
		fmt.Println(val)
	}

	// var stdoutBuff bytes.Buffer
	// defer stdoutBuff.WriteTo(os.Stdout)
	// intStream := make(chan int, 4)

	// go func() {
	// 	defer close(intStream)
	// 	defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
	// 		intStream <- i
	// 	}
	// }()

	// for integer := range intStream {
	// 	fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	// }

	go func() {
		var timer = time.Tick(2 * time.Second)

		for {
			select {
			case <-timer:
				fmt.Println("Tick")
			}
		}

	}()

	time.Sleep(2 * time.Second)

	var age Age

	age.AddOne()

	fmt.Println(age)

}

// Age ...
type Age int32

// AddOne adds one
func (a *Age) AddOne() {
	(*a)++
}
