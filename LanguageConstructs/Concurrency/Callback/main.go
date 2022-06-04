package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

type Worker struct{}

type callback func(io.Reader) error

// method compute will invoke callback after main processing is done
func (w *Worker) Compute(ctx context.Context, cb callback) {
	var buffer = bytes.NewBufferString("Done!")
	time.Sleep(1 * time.Second)

	// invoke passed in callback; which extends computation with custom logic
	cb(buffer)
}

func main() {
	var (
		worker Worker
		wait   = make(chan struct{})
	)

	worker.Compute(context.Background(), func(r io.Reader) error {
		defer close(wait)

		buffer, err := ioutil.ReadAll(r)
		fmt.Println(string(buffer))
		return err
	})

	<-wait
}
