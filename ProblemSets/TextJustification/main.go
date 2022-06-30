package main

import (
	"fmt"
	"math"
)

/*
Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.


Example 1:

Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Example 2:

Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
Note that the second line is also left-justified becase it contains only one word.
Example 3:

Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]
*/
func fullJustify(words []string, maxWidth int) []string {

	var (
		solution    []string
		cache       []string
		letterCount int
	)
	// each sequence of words that we pick for a line needs to have at least one space between them
	// Possible Solution 1:
	// We iterate over the words list and get a count of the chars in the word ~ assuming ascii
	// once words are chosen that are < maxWidth, we then

	flush := func(strs []string, maxWidth int) string {

		var (
			newStr      string
			spacesStart int
			charCount   int
		)

		for _, str := range strs {
			charCount += len(str)
		}

		spacesStart = maxWidth - charCount

		for idx, str := range strs {

			thisSpace := math.Ceil((float64(spacesStart) / float64(len(strs)-1-idx)))

			newStr += str

			for i := 0; i < int(thisSpace); i++ {
				newStr += " "
			}

			spacesStart -= int(thisSpace)

		}

		return newStr

	}

	for _, word := range words {

		// add word to cache iff len(word) + 1 + letterCount < maxWidth

		if len(word)+1+letterCount < maxWidth {
			letterCount += len(word) + 1
			cache = append(cache, word)
		} else {
			solution = append(solution, flush(cache, maxWidth))
			cache = []string{word}
			letterCount = len(word)
		}
	}

	solution = append(solution, flush(cache, maxWidth))

	return solution
}

func main() {

	var (
		words     []string
		justified []string
	)

	print := func(justified []string) {
		for _, val := range justified {
			fmt.Println(val)
		}
	}

	words = []string{"This", "is", "an", "example", "of", "text", "justification."}
	justified = fullJustify(words, 16)
	print(justified)

	fmt.Println("")

	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	justified = fullJustify(words, 16)
	print(justified)

	fmt.Println("")

	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	justified = fullJustify(words, 20)
	print(justified)

}
