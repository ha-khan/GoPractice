package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	anagrams := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(anagrams)

}

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
*/

func groupAnagrams(strs []string) [][]string {

	var (
		anagrams = make(map[string][]string)
		solution [][]string
	)

	for _, word := range strs {

		var (
			charStringArray []string
			sortedString    string
		)

		for _, char := range word {
			charStringArray = append(charStringArray, string(char))
		}

		sort.Strings(charStringArray)

		sortedString = strings.Join(charStringArray, "")

		anagrams[sortedString] = append(anagrams[sortedString], word)

	}

	for _, val := range anagrams {
		solution = append(solution, val)
	}

	return solution
}
