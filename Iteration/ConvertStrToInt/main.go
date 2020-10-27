package main

import (
	"fmt"
	"reflect"
	"time"
)

func yo(input ...int) {

	fmt.Println(len(input))

}

type (
	yochan chan int
)

var dude int

func (y *yochan) IsYo() {

	fmt.Println("erpiojtpyej")

}

type dope interface {
	IsYo()
}

var newErr error

type fn func() bool

type fnn fn

func (f *fn) What() {
	defer func() {

		if p := recover(); p != nil {
			fmt.Println("Recovered")
			fmt.Println(p)
		}
	}()

	panic(231)
}

func main() {

	var (
		cc        yochan
		cool, bro fn
		coool     fnn
		age       int = 23
		d         chan dope
	)

	fmt.Println(age)

	d = make(chan dope, 2)

	d <- &cc

	cc.IsYo()

	dude = 123

	cool = func() bool { return true }

	bro = cool

	bro()

	coool = func() bool { return true }

	go cool.What()

	time.Sleep(1 * time.Second)

	fmt.Println(string('a'))

	cool()
	coool()

	//asciiToInt := map[rune]int{rune()}
	y := "ȹsȹ"
	x := y[0]

	rx := reflect.TypeOf(x)
	fmt.Println(string(x))
	fmt.Println(rx.Size())

	for _, rn := range y {

		temp := reflect.TypeOf(rn)
		fmt.Println(temp.Name())
		fmt.Println(string(rn))
	}

	fmt.Println(4_2)

	mm := []int{1, 23, 4, 5, 5}

	fmt.Println(mm[-1])
}
