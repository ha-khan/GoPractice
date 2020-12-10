package main

import (
	"fmt"
	"sync"
)

/*
Suppose we have a class:

public class Foo {
  public void first() { print("first"); }
  public void second() { print("second"); }
  public void third() { print("third"); }
}

The same instance of Foo will be passed to three different threads.
Thread A will call first(), thread B will call second(), and thread C
will call third(). Design a mechanism and modify the program to ensure
that second() is executed after first(), and third() is executed after second().



Example 1:
Input: [1,2,3]
Output: "firstsecondthird"
Explanation: There are three threads being fired asynchronously.
The input [1,2,3] means thread A calls first(), thread B calls second(),
and thread C calls third(). "firstsecondthird" is the correct output.


Example 2:
Input: [1,3,2]
Output: "firstsecondthird"
Explanation: The input [1,3,2] means thread A calls first(),
thread B calls third(), and thread C calls second().
"firstsecondthird" is the correct output.


Note:

We do not know how the threads will be scheduled in the operating system,
even though the numbers in the input seems to imply the ordering. The input
format you see is mainly to ensure our tests' comprehensiveness.
*/

func newFoo() *foo {
	var (
		f = &foo{ /*l2: &sync.RWMutex{}, l3: &sync.RWMutex{}*/ }
	)

	f.l2 = new(sync.RWMutex) //&sync.RWMutex{}
	f.l3 = new(sync.RWMutex) //&sync.RWMutex{}

	f.l2.Lock()
	f.l3.Lock()

	return f
}

// potentially can also have a single variable such as an integer, where second is looping forever for value 2, and third is for 3
// the variable is initialized so that first can always execute.
// this allows the code to stay in userspace more often as

// locks are bit more overkill, but can reason about the problem easier as they are concurrency primitives
type foo struct {
	l2, l3 *sync.RWMutex
}

func (f *foo) first() {
	defer wg.Done()
	fmt.Print("first")
	f.l2.Unlock()
}

func (f *foo) second() {
	defer wg.Done()
	f.l2.Lock()
	fmt.Print("second")
	f.l3.Unlock()
}

func (f *foo) third() {
	defer wg.Done()
	f.l3.Lock()
	fmt.Print("third")
}

var wg sync.WaitGroup

func main() {

	f := newFoo()

	wg.Add(3)

	go f.first()
	go f.second()
	go f.third()

	wg.Wait()

}
