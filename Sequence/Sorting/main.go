package main

import (
	"GoPractice/utils"
	"fmt"
	"sort"
)

func main() {
	// slice := []int{2, 3, 5, 5, 5, 5, 5, 5, 6, 7, 8, 11, 11, 11, 13}
	slice := utils.BuildIntSlice(20, true)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//slice = mergesort(slice)
	//slice = quicksort(slice
	sort.Ints(slice) // NOTE: Within function call this will type cast to IntSlice()
	fmt.Println("\n--- Sorted ---\n\n", slice)
	deleteDuplicate(slice)
	fmt.Println("\n--- Delete Duplicate---\n\n", slice)
}

/*
Write a program which takes as input a sorted array and updates it so that
all duplicates have been removed and the remaining elements have been shifted
left to fill the emptied indices.
*/
func deleteDuplicate(data []int) {

	for i := 0; i < len(data) && i+1 < len(data); i++ {

		j := i + 1

		if data[i] == data[j] {
			for ; j < len(data); j++ {
				if data[i] == data[j] {
					data[j] = 0
				} else {

					// create copy slice of 0s
					zero := make([]int, len(data[i+1:j]))
					copy(zero, data[i+1:j])

					// reslice to remove 0s
					data = append(data[:i+1], data[j:]...)
					data = append(data, zero...)

					break
				}

			}
		}

	}

}

func mergesort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	midPoint := len(a) / 2

	left := mergesort(a[:midPoint])
	right := mergesort(a[midPoint:])

	return merge(left, right)
}

func merge(a, b []int) []int {

	sorted := make([]int, 0)

	idxA := 0
	idxB := 0

	for idxA < len(a) && idxB < len(b) {
		if a[idxA] < b[idxB] {
			sorted = append(sorted, a[idxA])
			idxA++
		} else {
			sorted = append(sorted, b[idxB])
			idxB++
		}
	}

	// Add remaining
	for ; idxA < len(a); idxA++ {
		sorted = append(sorted, a[idxA])
	}

	for ; idxB < len(b); idxB++ {
		sorted = append(sorted, b[idxB])
	}

	return sorted
}

func quicksortmemory(list []int, pivot int) []int {

	var (
		lessThan, greaterThan, sorted []int
		pivotVal                      int
	)

	if len(list) == 0 {
		return list
	}

	pivotVal = list[pivot]

	for _, val := range list {
		if val < pivotVal {
			lessThan = append(lessThan, val)
		} else {
			greaterThan = append(greaterThan, val)
		}
	}

	// Recurse
	sortedLessThan := quicksortmemory(lessThan, len(lessThan)/2)
	sortedGreaterThan := quicksortmemory(greaterThan, len(greaterThan)/2)

	// build up sorted
	for _, val := range sortedLessThan {
		sorted = append(sorted, val)
	}
	for _, val := range sortedGreaterThan {
		sorted = append(sorted, val)
	}

	return sorted
}
