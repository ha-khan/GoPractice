package main

import (
	"fmt"
	"math"
)

/*
Given a string of symbols A and B, goal is to convert it to the form A...AB...B

so all A symbols need to occur before B symbols.


goal is to return the minimum number of letters that need to be deleted from s in order to obtain a string in the form
A...AB...B


s = BAAABAB -> AAABB ret 2
s = BBABAA -> ret 3
s = AABBBB -> ret 0
*/

// Linearly scan resulting string and if it is well formed A...AB...B
// then return that it is correct
func checkWellFormed(s string) bool {

	context := "A"
	change := false

	for idx := range s {
		if string(s[idx]) == context {

		} else if change {
			return false
		} else {
			context = "B"
			change = true
		}
	}

	return true
}

// Iterate over passed in string, and recurse with two calls:
// Remove symbol at idx or not remove at idx add + 1 for each removal and increment idx to next start point..
// The resulting string (end node) should then be checked, and if its not in the correct form then add a weight of 10000000
// TODO: Optimize by stopping recursion prematurely and probably cache results also
// Essentially a DFS on the recursive call graph
func sortTwoWord(s string, idx int) float64 {

	// Base case is when idx is incremented beyond size of string
	if idx > len(s)-1 {

		if checkWellFormed(s) {
			return float64(0)
		}

		return float64(1000000)
	}

	return math.Min(sortTwoWord(s[:], idx+1), float64(1)+sortTwoWord(s[:idx]+s[idx+1:], idx))
}

func main() {
	fmt.Println(sortTwoWord("BAAABAB", 0))
	fmt.Println(sortTwoWord("BBABAA", 0))
	fmt.Println(sortTwoWord("AABBBB", 0))
	fmt.Println(sortTwoWord("BAAABABAABBBA", 0))
}
