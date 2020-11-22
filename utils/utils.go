package utils

import (
	"math/rand"
	"time"
)

// CopyMap returns a copy of the MAP passed in
func CopyMap(env map[string]interface{}) map[string]interface{} {

	var mp = make(map[string]interface{})

	for key, val := range env {
		mp[key] = val
	}

	return mp

}

// Min returns the minimum of the two integer params
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Max returns the maximum of the two integer params
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

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
