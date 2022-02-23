package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var notify chan struct{}
	notify = make(chan struct{})
	ctx, cancel := context.WithCancel(context.TODO())
	go StartApp(ctx, notify)
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	<-notify

}

func StartApp(ctx interface{}, notify chan struct{}) {
	fmt.Println("starting...")
	switch t := ctx.(type) {
	case context.Context:
		<-t.Done()
		notify <- struct{}{}
	default:
		panic("Unknown type!")
	}
}
