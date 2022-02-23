package main

import (
	"fmt"
)

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:

Input: nums = [1]
Output: 1

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23


*/
func main() {

	nums := []int{5, 4, -1, 7, 8}

	fmt.Println(fmt.Sprintf("Largest nums %+v is %+v", nums, maximumSubarraySumSlow(nums)))

}

func sum(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

/*
Two pointers iterating redundantly

space ~ O(1)

time ~ O(n^2)
*/
func maximumSubarraySumSlow(nums []int) int {
	var largestSum int
	for idx := 0; idx < len(nums); idx++ {
		for idy := idx; idy < len(nums); idy++ {
			if total := sum(nums[idx : idy+1]); largestSum < total {
				largestSum = total
			}
		}
	}

	return largestSum
}

/*
dynamic programming

space ~ O(n)

time ~ O(n)


func maximumSubarraySum(nums []int) int {
	var (
		cache = make(map[string]int)
	)

	return 0
}
*/
