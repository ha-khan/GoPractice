package utils

import (
	"math/rand"
	"time"
)

// BuildIntSlice of random integers
func BuildIntSlice(size int, positive bool) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if !positive {
			slice[i] = rand.Intn(999) - rand.Intn(999)
		} else {
			slice[i] = rand.Intn(999)
		}
	}
	return slice
}

// BuildMatrix returns an MxN matrix or a slice of slice of ints
func BuildMatrix(m, n int) [][]int {

	var matrix [][]int

	for idx := 0; idx < m; idx++ {
		matrix = append(matrix, BuildIntSlice(n, true))
	}

	return matrix
}
