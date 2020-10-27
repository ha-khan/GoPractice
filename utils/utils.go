package utils

import (
	"math/rand"
	"time"
)

// Yo ....
const Yo = 1000

// GenerateSlice of random integers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
