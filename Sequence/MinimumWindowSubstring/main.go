package main

import (
	"fmt"
)

func main() {
	var (
		s, t string
	)

	s = "ADOBECODEBANC"
	t = "ABC"
	fmt.Println(minWindow(s, t))
	fmt.Println()
	s = "a"
	t = "a"
	fmt.Println(minWindow(s, t))
}

/*
Given two strings s and t, return the minimum window in s which will contain all the characters in t. If there is no such window in s that covers all characters in t, return the empty string "".

Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Example 2:

Input: s = "a", t = "a"
Output: "a"


Constraints:

1 <= s.length, t.length <= 105
s and t consist of English letters.


Follow up: Could you find an algorithm that runs in O(n) time?
*/
type counter map[rune]int

func (c *counter) add(r rune) {

	if _, ok := (*c)[r]; ok {
		(*c)[r]++
	} else {
		(*c)[r] = 1
	}
}

func (c *counter) remove(r rune) {

	if _, ok := (*c)[r]; ok {
		(*c)[r]--
		if (*c)[r] < 0 {
			delete(*c, r)
		}
	}
}

func (c *counter) member(r rune) bool {
	_, ok := (*c)[r]
	return ok
}

func (c *counter) count() int {
	var (
		count int
	)

	for _, val := range *c {
		count += val
	}

	return count
}

func (c *counter) copy() *counter {
	var (
		cpy = make(counter)
	)

	for key, val := range *c {
		cpy[key] = val
	}

	return &cpy
}

func minWindow(s string, t string) string {
	var (
		ctr       = make(counter)
		minSubStr = s
		found     = false
	)

	// Count char/rune(s) of string t
	for _, val := range t {
		ctr.add(val)
	}

	for left := 0; left < len(s); {

		char := rune(s[left])
		if ctr.member(char) {
			right := left

			cpy := ctr.copy()
			// Get a copy of the counter
			// for each char at index right, we decrement the values
			// we also check each time if the count() of the counter is 0, implying that all chars have been hit
			for ; right < len(s); right++ {
				cpy.remove(rune(s[right]))
				if cpy.count() == 0 && len(s[left:right+1]) <= len(minSubStr) {
					minSubStr = s[left : right+1]
					found = true
				}
			}
		}

		// minSubStr contains the smallest window with all the characters from the start left index
		// now we move left over to the very next char that belongs in ctr
		left++
		for left < len(s) {
			char := rune(s[left])
			if !ctr.member(char) {
				left++
			} else {
				break
			}
		}
	}

	if found {
		return minSubStr
	}

	return ""
}
