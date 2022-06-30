package main

import (
	"fmt"
)

// O(n^2)
func subStringNoRepeat(str string) string {

	var longest string

	for idx := range str {
		mp := make(map[rune]interface{})
		var substr string

		for _, idy := range str[idx:] {
			if _, ok := mp[rune(idy)]; ok {
				break
			}

			mp[rune(idy)] = true
			substr += string(idy)
		}

		if len(substr) > len(longest) {
			longest = substr
		}
	}

	return longest
}

// O(n)
func subStringNoRepeatFast(str string) string {

	var (

		// cache holds chars seen between str[startIdx:idx], and is used to check for
		// uniqueness, works like a Set as the Key is the only input that matters
		cache = make(map[string]interface{})

		// holds the longest substr seen thus far w/o repeating chars
		substr string

		// dynamic sliding window technique where startIdx:idx should contain all unique chars
		// and if we collide when incrementing the idx, then we move startIdx++ until the colliding chars are gone
		// Everytime we increment idx and the char is unique, we do a check to see if the str[startIdx:idx+1] is larger than
		// the currently stored substring
		startIdx int
	)

	for idx := 0; idx < len(str); idx++ {

		if _, ok := cache[string(str[idx])]; !ok {
			cache[string(str[idx])] = true
		} else {
			// if we did collide with a char, then need to increment startIdx..
			for startIdx < idx {
				// remove the char that startIdx is pointing to in the cache
				// check whether the collision still holds with idx,
				// do until no collision.
				delete(cache, string(str[startIdx]))
				startIdx++
				if _, ok := cache[string(str[idx])]; !ok {
					cache[string(str[idx])] = true
					break
				}
			}
		}

		// Check if longest substr seen thus far
		if idx+1 < len(str) {
			if len(substr) < len(str[startIdx:idx+1]) {
				substr = str[startIdx : idx+1]
			}
		}
	}

	return substr
}

func main() {

	fmt.Println(subStringNoRepeat("abcabcbb"))
	fmt.Println(subStringNoRepeat("bbbbbbb"))
	fmt.Println(subStringNoRepeatFast("abcabcbb"))
	fmt.Println(subStringNoRepeatFast("bbbbbb"))
	fmt.Println(subStringNoRepeatFast("pwwkew"))
	fmt.Println(subStringNoRepeatFast("pwwkewehgpmhazsaetqp"))

}
