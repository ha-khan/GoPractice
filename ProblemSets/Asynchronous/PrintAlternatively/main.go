package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
The same instance of FooBar will be passed to two different threads.

Thread A will call foo() while thread B will call bar().

Modify the given program to output "foobar" n times.

Example 1:

Input: n = 1
Output: "foobar"

Explanation: There are two threads being fired asynchronously.

One of them calls foo(), while the other calls bar().

"foobar" is being output 1 time.

Example 2:
Input: n = 2
Output: "foobarfoobar"
Explanation: "foobar" is being output 2 times.
*/
func main() {
	var err error
	var n int
	var scanner = bufio.NewScanner(os.Stdin)

read:
	fmt.Print("Please input an integer: ")
	scanner.Scan()

	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		goto read
	}

	f := newFooBar(n)
	wg.Add(2)

	go f.foo()
	go f.bar()

	wg.Wait()
	fmt.Println(f)
}

var wg sync.WaitGroup

func newFooBar(in int) *foobar {
	var fb = &foobar{
		n:       in,
		fooLock: new(sync.RWMutex),
		barLock: new(sync.RWMutex),
		Builder: new(strings.Builder),
	}

	// initially lock the bar go routine
	// which is unlocked only when
	fb.barLock.Lock()

	return fb
}

type foobar struct {
	n                int
	fooLock, barLock *sync.RWMutex

	// buffer to store string that we're building
	*strings.Builder
}

func (f *foobar) foo() {
	defer wg.Done()
	for idx := 0; idx < f.n; idx++ {
		f.fooLock.Lock()
		f.WriteString("foo")
		f.barLock.Unlock()
	}
}

func (f *foobar) bar() {
	defer wg.Done()
	for idx := 0; idx < f.n; idx++ {
		f.barLock.Lock()
		f.WriteString("bar")
		f.fooLock.Unlock()
	}
}
