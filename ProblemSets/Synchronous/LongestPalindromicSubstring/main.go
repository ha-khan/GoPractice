package main

import (
	"fmt"
)

// Approach 1:
// 		Compute/Enumerate all possible substrings (brute force) and check if they are palindromes
// 		and keep track of largest seen palindrome thus far. This is assuming we have no information
// 		that we can work off of/utilize to shorten the "search"
//
//		O(n^3)
func longestPalindromicSubStringSlow(str string) string {

	var (
		max string
	)

	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if isMax := checkIfPalindromeAndGetSize(str[i : j+1]); len(isMax) > len(max) {
				max = isMax
			}
		}
	}

	return max
}

/*
	should be around O(mn), though in the case where the string is the same set of chars
	aaaaaaaaaa, algorithm will not have the greatest performance O(n^2)
*/
func longestPalindromicSubStringFast(str string) string {

	var (
		// Keeps track of indices where chars/runes appear
		cache = make(map[rune][]int)

		max string
	)

	for idx := 0; idx < len(str); idx++ {
		char := rune(str[idx])
		if val, ok := cache[char]; !ok {
			cache[char] = append(cache[char], idx)
		} else {
			for _, idy := range val {
				if isMax := checkIfPalindromeAndGetSize(str[idy : idx+1]); len(isMax) > len(max) {
					max = isMax
				}
			}
		}

	}

	return max

}

// Helper function to do quick check
func checkIfPalindromeAndGetSize(str string) string {

	i := 0
	j := len(str) - 1

	for i < j {

		if str[i] != str[j] {
			return ""
		}

		i++
		j--
	}
	return str

}

func main() {
	fmt.Println(longestPalindromicSubStringFast("babad"))
	fmt.Println(longestPalindromicSubStringFast("racecar"))
	fmt.Println(longestPalindromicSubStringFast("cbbd"))
	fmt.Println(longestPalindromicSubStringFast("abcdefghtacocatijklol"))
}
