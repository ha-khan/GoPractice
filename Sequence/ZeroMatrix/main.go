package main

import (
	"GoPractice/utils"
	"fmt"
)

func zeroMatrix(matrix ...[]int) [][]int {

	// Store columns which have been zeroed out
	var cache = make(map[int]interface{})
	var rowSize = len(matrix[0])

	fn := func(col int, matrix ...[]int) {
		for idx := range matrix {
			matrix[idx][col] = 0
		}
	}

	for idx := range matrix {
		for idy := range matrix[idx] {
			if matrix[idx][idy] != 0 {
				continue
			}
			if _, ok := cache[idy]; !ok {
				matrix[idx] = make([]int, rowSize)
				fn(idy, matrix...)
				cache[idy] = true
				break
			}
		}
	}

	return matrix
}

func main() {

	var matrix = utils.BuildMatrix(5, 4)

	matrix[2][3] = 0
	matrix[0][0] = 0

	matrix = zeroMatrix(matrix...)

	for _, val := range matrix {
		fmt.Println(val)
	}

}
