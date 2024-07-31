package main

import (
	"log"
	"math/rand"
	"slices"
)

func main() {
	var data []int
	for n := range 0xFF {
		data = append(data, n)
	}

	target := rand.Intn(len(data))
	loc, ok := slices.BinarySearch(data, target)
	if !ok {
		log.Fatalf("list does not contain target=%d", target)
	}
	log.Printf("target '%d' is at index '%d'", target, loc)

	var queue = make(chan int, len(data))
	iterator := slices.All(data)
	for _, val := range iterator {
		go func() {
			queue <- val
		}()
	}

loop:
	if len(queue) != len(data) {
		goto loop
	}

	close(queue)
	for val := range queue {
		log.Printf("val = %d\n", val)
	}
}
