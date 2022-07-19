package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"sync"
	"time"
)

type Worker struct{}

type callback func(io.Reader, error) error

// method compute will invoke callback after main processing is done
func (w *Worker) Compute(ctx context.Context, cb callback) {
	var buffer = bytes.NewBufferString("Done!")
	time.Sleep(1 * time.Second)

	// invoke passed in callback; which extends computation with custom logic
	cb(buffer, nil)
}

func main() {
	var worker Worker
	var wg = sync.WaitGroup{}

	wg.Add(1)
	worker.Compute(context.Background(), func(r io.Reader, e error) error {
		defer wg.Done()

		buffer, err := ioutil.ReadAll(r)
		fmt.Println(string(buffer))
		return err
	})

	wg.Wait()
}
