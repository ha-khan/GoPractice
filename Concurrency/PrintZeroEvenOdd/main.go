package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

/*
You have a function printNumber that can be called with an integer parameter and prints it to the console.

    For example, calling printNumber(7) prints 7 to the console.

You are given an instance of the class ZeroEvenOdd that has three functions: zero, even, and odd.

The same instance of ZeroEvenOdd will be passed to three different threads:

    Thread A: calls zero() that should only output 0's.
    Thread B: calls even() that should only output even numbers.
    Thread C: calls odd() that should only output odd numbers.

Modify the given class to output the series "010203040506..." where the length of the series must be 2n.

Implement the ZeroEvenOdd class:

    ZeroEvenOdd(int n) Initializes the object with the number n that represents the numbers that should be printed.
    void zero(printNumber) Calls printNumber to output one zero.
    void even(printNumber) Calls printNumber to output one even number.
    void odd(printNumber) Calls printNumber to output one odd number.



Example 1:

Input: n = 2
Output: "0102"
Explanation: There are three threads being fired asynchronously.
One of them calls zero(), the other calls even(), and the last one calls odd().
"0102" is the correct output.

Example 2:

Input: n = 5
Output: "0102030405"



Constraints:

    1 <= n <= 1000

*/
func main() {
	var (
		wg      = &sync.WaitGroup{}
		scanner = bufio.NewScanner(os.Stdin)
	)

	printNumber := print(func(n int) {
		fmt.Print(n)
		if n != 0 {
			wg.Done()
		}
	})

	fmt.Print("Input an integer: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err.Error())
	}

	controller := NewZeroEvenOdd(n)

	wg.Add(n)

	go controller.Zero(printNumber)
	go controller.Odd(printNumber)
	go controller.Even(printNumber)

	controller.zero <- controller.odd

	wg.Wait()
	close(controller.zero)
	close(controller.even)
	close(controller.odd)
}

func driver(n int) string {
	var (
		wg  = &sync.WaitGroup{}
		str string
	)

	printNumber := print(func(n int) {
		str = (str + strconv.Itoa(n))
		if n != 0 {
			wg.Done()
		}
	})

	controller := NewZeroEvenOdd(n)

	wg.Add(n)

	go controller.Zero(printNumber)
	go controller.Odd(printNumber)
	go controller.Even(printNumber)

	controller.zero <- controller.odd

	wg.Wait()
	close(controller.zero)
	close(controller.even)
	close(controller.odd)

	return str
}

type print func(int)

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	return &ZeroEvenOdd{n: n, zero: make(chan chan struct{}), even: make(chan struct{}), odd: make(chan struct{})}
}

type ZeroEvenOdd struct {
	n         int
	zero      chan chan struct{} // bi-directional chan that recvs send only chan
	even, odd chan struct{}
}

// need to keep count that last
func (z *ZeroEvenOdd) Zero(fn print) {
	for writeChan := range z.zero {
		fn(0)
		writeChan <- struct{}{}
	}
}

func (z *ZeroEvenOdd) Even(fn print) {
	for i := 2; i <= z.n; i++ {
		if i%2 == 0 {
			<-z.even
			fn(i)
			// need to skip this final send if i == n which is even, so reached
			if !(i == z.n && (z.n%2 == 0)) {
				z.zero <- z.odd
			}
		}
	}
}

func (z *ZeroEvenOdd) Odd(fn print) {
	for i := 1; i <= z.n; i++ {
		if i%2 != 0 {
			<-z.odd
			fn(i)
			// need to skip this final send if i == n which is odd, so reached
			if !(i == z.n && (z.n%2 != 0)) {
				z.zero <- z.even
			}
		}
	}
}
