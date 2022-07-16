package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	goal is to create a new DatabaseConnector if there isn't one
	already allocated.

	if no in-flight requests are using the DatabaseConnector, then Close() it

	if another request happens, need to recreate the DatabaseConnector by calling
	the NewDatabaseConnector constructor func

	need to also limit amount of calls that can be done on the Worker.Execute(..) method
	as we have a constraint on the amount of calls that can be done on a single instance of
	the DatabaseConnector, each extra call on Execute should block the caller, and eventually
	should unblock if "room" is available to call Worker.Execute

	assume that the

*/
func main() {
	var w = new(Worker)
	w.Semaphore = NewSemaphore()

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			w.Execute(fmt.Sprintf("Query %v", i))
			wg.Done()
		}()
	}

	wg.Wait()
}

type Worker struct {
	// single reference to a db client
	DBClient *DatabaseConnector

	// counter to keep track of currently processing requests
	// a Lock is also placed on it as there is a potential race
	// where the counter can be 0 and multiple go routines can view that
	Counter     int
	CounterLock sync.RWMutex

	// keeps track of total amount of 'in-flight' requests
	Semaphore *Semaphore
}

// not thread safe what would
func (w *Worker) Execute(args string) {
	w.Semaphore.Acquire()
	defer w.Semaphore.Release()

	w.CounterLock.Lock()
	if w.Counter == 0 {
		w.DBClient = NewDatabaseConnector()
		w.Counter++
	}
	w.CounterLock.Unlock()

	w.DBClient.ExecuteQuery(fmt.Sprintf("db.select.%s", args))

	w.CounterLock.Lock()
	w.Counter--
	if w.Counter == 0 {
		w.DBClient.Close()
	}
	w.CounterLock.Unlock()
}

type Semaphore struct {
	queue chan struct{}
}

func NewSemaphore() *Semaphore {
	return &Semaphore{
		queue: make(chan struct{}, 1),
	}
}

func (s *Semaphore) Acquire() {
	s.queue <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.queue
}

//
//
//
func NewDatabaseConnector() *DatabaseConnector {
	fmt.Println("Creating DatabaseConnector")
	time.Sleep(time.Second * 1)
	return new(DatabaseConnector)
}

type DatabaseConnector struct{}

// method is considered thread safe, so can be called by multiple go routines at once
func (d *DatabaseConnector) ExecuteQuery(query string) error {
	fmt.Printf("Executing: %s\n", query)
	return nil
}

// closes the DatabaseConnector by releasing any lower level resources
func (d *DatabaseConnector) Close() error {
	fmt.Println("Closing DatabaseConnector")
	time.Sleep(time.Second * 1)
	return nil
}
