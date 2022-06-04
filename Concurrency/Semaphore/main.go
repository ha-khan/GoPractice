package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// semaphore locks computation to at most 5 running concurrently
var sem = semaphore.NewWeighted(5)

// weight group keeps track of all executing computations
// when decremented to 0, will close errorQueue channel
var wg = sync.WaitGroup{}

// rw channel
// send/receive channel
type errorQueue chan error

type Worker struct {
	eq errorQueue
}

type Dispatcher struct {
}

type Buffer interface {
	Write([]byte) (int, error)
}

// thread safe
func (w *Worker) Computation(data []byte, b Buffer) {
	sem.Acquire(context.Background(), 1)
	defer sem.Release(1)
	defer wg.Done()
	time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)
	fmt.Println("COMPUTATION")
}

func main() {
	var (
		eq = make(errorQueue)
		w  = Worker{eq: eq}
		i  = 0
	)

	// go routine keeps track of all computations
	// will close channel once all signal complete => wg at 0
	go func() {
		wg.Wait()
		close(eq)
	}()

	wg.Add(20)
	for i < 20 {
		go w.Computation(nil, nil)
		i++
	}

	// reads until chan closed
	for err := range eq {
		log.Printf("error: %s", err.Error())
	}
}
