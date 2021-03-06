package main

import (
	"fmt"
)

/*
Given two words (beginWord and endWord), and a dictionary's word list,
find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

*/

func main() {

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	shortestLength := ladderLength(beginWord, endWord, wordList)
	fmt.Println(shortestLength)

	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	shortestLength = ladderLength(beginWord, endWord, wordList)
	fmt.Println(shortestLength)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	var (
		// Set used to quickly check whether newly created words are valid.
		// i.e. if they are in given wordList
		validWords = make(map[string]interface{})

		// Each idx/char of a word can be transformed to another word iff the wordList
		// contains a word with aa char at that idx, therefore need to process the wordList and at
		// each idx come up with all possible char transformations
		validTransf = make([]map[rune]interface{}, len(beginWord))
	)

	// Populate set of valid words
	for _, val := range wordList {
		validWords[val] = true
	}

	// Populate set of valid transformations for each idx of beginWord
	for _, val := range wordList {
		for idx, char := range val {

			// skip char at position as its not a valid transformation
			if char == rune(beginWord[idx]) {
				continue
			}

			if validTransf[idx] == nil {
				validTransf[idx] = make(map[rune]interface{})
			}
			validTransf[idx][char] = true
		}
	}

	// Do a BFS for each valid and unique transformation from beginWord, which is a graph.
	// The graph will not be pre computed but will be constructed at runtime, word by word.
	// BFS guarantees that we will get the shortest path between two nodes, so the shorted path/transformations from
	// beginWord -> endWord where we should termiante, or if we are unable to transform any further, then stop
	alreadySeen := make(map[string]interface{})
	queue := []string{beginWord}
	counter := 1
	for {
		var (
			queue2 []string
		)

		// If we reach the end of the queue with w/o returning, then transformations to
		// endWord does not exist
		if len(queue) == 0 {
			counter = 0
			break
		}

		counter++

		for _, word := range queue {
			for idx, chars := range validTransf {
				for char := range chars {
					var (
						newWord string
					)

					// Need to replace the char at index idx with char
					// Recall that word[idx+1:] if idx == len(word) will be an empty string ""
					newWord = word[:idx] + string(char) + word[idx+1:]

					if newWord == endWord {
						return counter
					}

					if _, ok := validWords[newWord]; ok {
						if _, ok = alreadySeen[newWord]; !ok {
							queue2 = append(queue2, newWord)
							alreadySeen[newWord] = true
						}
					}
				}
			}
		}

		queue = queue2
	}

	return counter
}
