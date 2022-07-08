package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

func main() {
	// set runtime state
	// setting to 1 forces a single OS thread to be
	// scheduled work specified by some go routine
	// OS itself would schedule the oS thread to whatever
	// available core
	//
	// setting to 1 "slows" the sorting as we're making the program
	// sequential, even if we're launching 2 go routines
	fmt.Println(runtime.GOMAXPROCS(2))
	rand.Seed(time.Now().Unix())

	var start = time.Now()
	var done = make(chan []int, 2)
	defer close(done)

	var size = 10_000_000
	var halfSize = size / 2

	var buffer = make([]int, size)
	for i := 0; i < size; i++ {
		buffer[i] = rand.Int()
	}

	fn := func(ints []int) {
		sort.Sort(sort.IntSlice(ints))
		done <- ints
	}

	go fn(buffer[:halfSize])
	go fn(buffer[halfSize:])

	<-done
	<-done

	// sort entire buffer as each half itself has been sorted
	// independent of one another
	sort.Sort(sort.IntSlice(buffer))
	fmt.Println(time.Now().Sub(start))
}
