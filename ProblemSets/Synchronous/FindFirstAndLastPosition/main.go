package main

import "fmt"

func main() {
	output := searchRange([]int{5, 7, 7, 8, 8, 10}, 5)
	fmt.Println(output)

	output = searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println(output)

	output = searchRange([]int{}, 0)
	fmt.Println(output)
}

/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

Follow up: Could you write an algorithm with O(log n) runtime complexity?



Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]

*/
func searchRange(nums []int, target int) []int {

	idx := binarySearchRecursion(nums, 0, len(nums)-1, target)

	if idx < 0 {
		return []int{-1, -1}
	}

	// Once we find the target value in the sorted sequence of number, we need to
	// expand "outward" from left & right until we reach numbers that are not the target & stop
	i, j := idx, idx

	for {
		if i >= 0 && nums[i] == target {
			i--
		} else {
			break
		}
	}

	for {
		if j < len(nums)-1 && nums[j] == target {
			j++
		} else {
			break
		}
	}

	return []int{i + 1, j - 1}

}

// Will return the idx of the number if exists,
// postive implies existence; negative implies not exist
func binarySearchRecursion(nums []int, low, high /*needs to be len(nums) - 1*/, target int) int {

	if low > high {
		//fmt.Println(fmt.Sprintf("low: %d, high: %d", low, high))
		return -1
	}

	// Go will auto take the floor
	midpoint := (high + low) / 2

	if nums[midpoint] < target {
		return binarySearchRecursion(nums, midpoint+1, high, target)
	} else if nums[midpoint] > target {
		return binarySearchRecursion(nums, low, midpoint-1, target)
	} else {
		return midpoint
	}

}
