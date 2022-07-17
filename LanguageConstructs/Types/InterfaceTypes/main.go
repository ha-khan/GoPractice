package main

import (
	"fmt"
	"sync"
)

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore chan struct{}

func (s *semaphore) Acquire() {
	*s <- struct{}{}
}

func (s *semaphore) Release() {
	<-*s
}

type LifoQueue interface {
}

func main() {
	var s = make(semaphore, 2)
	var wg = sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			s.Acquire()
			fmt.Println("Running")
			defer s.Release()
			defer wg.Done()
		}()
	}

	wg.Wait()

	// type assertion

	// type switch
}
