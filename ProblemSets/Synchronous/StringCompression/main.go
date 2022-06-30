package main

import (
	"fmt"
	"strconv"
)

/*

Write a function that "compresses" string str : aabcccccaaa --> a2b1c5a3

if "compressed" string is greater than uncompressed, then return uncompressed string


*/

// Assume A-Z, a-Z
/*

Approach is two iterators that utilize the dynamically resizing sliding window technique
where the slow iterator is the beginning of the substr and the fast will iterate and count until sequential chars/runes
are not found, in which case we then increment slow to fast and append the

*/
func compressString(str string) string {

	var (
		compressedStr string
		slow, fast    int
	)

	// Iterate slowly
	for slow = 0; slow < len(str) && fast < len(str); {
		var (
			currentChar rune
			counter     int
		)

		currentChar = rune(str[slow])

		// Iterate quickly
		for fast = slow; fast < len(str); fast++ {
			if currentChar != rune(str[fast]) {
				break
			}

			counter++
		}

		compressedStr += string(currentChar) + strconv.Itoa(counter)
		slow = fast
	}

	if len(compressedStr) < len(str) {
		return compressedStr
	}

	return str
}

func main() {

	str := compressString("aabcccccaaa")
	fmt.Println(str)

}
