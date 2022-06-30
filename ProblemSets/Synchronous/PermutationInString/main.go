package main

import (
	"fmt"
)

func main() {

	fmt.Println(permutationString("ab", "eidbaooo"))
	fmt.Println(permutationString("ab", "eidboaoo"))

}

/*
Given two strings s1 and s2, write a function to return true if s2 contains
the permutation of s1. In other words, one of the first string's permutations
is the substring of the second string

Example 1:
Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False

	With the passed in permutation string we need to create a mapping of it
	and then use the dynamic sliding window technique to check each possible substring of the string
	and ensure that it obeys the following conditions.

	1. Len is the same as the permutation
	2. All chars/Runes in the substring match

	Time ~ O(n)
	Space ~ O(n) b/c of cache...
*/
func permutationString(permutation, original string) bool {

	var (
		cache map[rune]interface{}
	)

	populateCache := func() map[rune]interface{} {
		var (
			cache = make(map[rune]interface{})
		)
		for _, val := range permutation {
			cache[val] = true
		}
		return cache
	}

	cache = populateCache()

	// As soon as we encounter a rune/char that exists in the cache, we check the next consequtive entries
	// if we get all of them then break and exit with true, else skip forward
	for idx := 0; idx < len(original); {
		for idy := idx; idy < len(original); {

			//Could also potentially remove modification of cache
			if _, found := cache[rune(original[idy])]; found {
				delete(cache, rune(original[idy]))
			} else {
				if len(cache) == 0 {
					return true
				}
				cache = populateCache()
			}
			idy++
			idx = idy
		}

	}

	return false
}
