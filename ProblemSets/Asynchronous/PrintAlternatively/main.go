package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func newFooBar(in int) *foobar {
	var ff = &foobar{n: in}

	ff.f = new(sync.RWMutex)
	ff.g = new(sync.RWMutex)

	ff.g.Lock()

	return ff
}

type foobar struct {
	n    int
	f, g *sync.RWMutex
	// add extra bool variable constant??
	// flag bool ~ 0 value is always false...
}

func (f *foobar) foo() {
	defer wg.Done()
	for idx := 0; idx < f.n; idx++ {
		f.f.Lock()
		fmt.Print("foo")
		f.g.Unlock()
	}
}

func (f *foobar) bar() {
	defer wg.Done()
	for idx := 0; idx < f.n; idx++ {
		f.g.Lock()
		fmt.Print("bar")
		f.f.Unlock()

	}
}

/*
The same instance of FooBar will be passed to two different threads.

Thread A will call foo() while thread B will call bar().

Modify the given program to output "foobar" n times.



Example 1:

Input: n = 1
Output: "foobar"

Explanation: There are two threads being fired asynchronously. One of them calls foo(), while the other calls bar(). "foobar" is being output 1 time.

Example 2:
Input: n = 2
Output: "foobarfoobar"
Explanation: "foobar" is being output 2 times.

*/
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please input an integer: ")

	// get input
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	f := newFooBar(n)

	wg.Add(2)

	go f.foo()
	go f.bar()

	wg.Wait()
}
