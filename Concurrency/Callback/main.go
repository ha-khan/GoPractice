package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"
)

var buffer = bytes.NewBufferString("")

type Computation struct{}

type callback func(io.Writer) error

// method compute will invoke callback after main processing is done
func (c *Computation) Compute(ctx context.Context, cb callback) {
	time.Sleep(1 * time.Second)
	cb(buffer)
}

func main() {
	var (
		worker Computation
		wait   = make(chan struct{})
	)

	worker.Compute(context.Background(), func(w io.Writer) error {
		defer close(wait)
		_, err := w.Write([]byte("done"))
		return err
	})

	<-wait
	fmt.Println(buffer)
}
