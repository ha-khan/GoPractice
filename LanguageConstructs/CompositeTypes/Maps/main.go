package main

import (
	"fmt"
	"sync"
	"time"
)

type executor func() error

func main() {
	var invoker = make(map[string]executor)

	invoker["traversal"] = executor(func() error {
		fmt.Println("Executing")
		return nil
	})

	if fn, ok := invoker["traversal"]; ok {
		fn()
	}

	var threadSafeMap = new(sync.Map)
	var done = make(chan struct{}, 0)
	threadSafeMap.Store("traversal", executor(func() error {
		fmt.Println("Executing")
		done <- struct{}{}
		return nil
	}))

	if fn, ok := threadSafeMap.Load("traversal"); ok {
		go fn.(executor)()
	}

	var time = time.Tick(time.Second * 5)
	select {
	case <-done:
		fmt.Println("Done")
	case <-time:
		fmt.Println("Timer ran out!")
	}
}
