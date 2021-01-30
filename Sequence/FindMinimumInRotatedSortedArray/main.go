/*
Given a rotated sorted array, find the minimum. Assume there are no duplicates

Input: [8, 9, 0, 1, 2, 3, 4]
Output: 0

Input: [2,3,4,5,1]
Output: 1

Input: [1,2,3,4]
Output: 1
*/

package main

import "fmt"

func main() {
	rotatedArray := []int{8, 9, 0, 1, 2, 3, 4}
	minimum := findMinInRotatedSortedArray(rotatedArray)
	fmt.Println(minimum)

	rotatedArray = []int{2, 3, 4, 5, 1}
	minimum = findMinInRotatedSortedArray(rotatedArray)
	fmt.Println(minimum)

	rotatedArray = []int{1, 2, 3, 4}
	minimum = findMinInRotatedSortedArray(rotatedArray)
	fmt.Println(minimum)
}

func findMinInRotatedSortedArray(array []int) int {

	var (
		minimum = 10000
	)

	idx, idy := 0, len(array)-1

	for idx <= idy {

		// takes floor if not int
		midPoint := (idx + idy) / 2

		entry := array[midPoint]

		if entry < minimum {
			minimum = entry
		}

		leftMost := array[idx]
		rightMost := array[idy]

		// Implies rotation at 0, so sorted array
		if leftMost < entry && rightMost > entry {
			idy = midPoint - 1
		} else if leftMost > entry {
			idy = midPoint - 1
		} else {
			idx = midPoint + 1
		}

	}

	return minimum
}

// Your last C/C++ code is saved below:
// #include <iostream>
// using namespace std;

// int main() {
// 	cout<<"Hello";
// 	return 0;
// }
