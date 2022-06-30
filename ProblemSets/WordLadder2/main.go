package main

import (
	"fmt"
)

/*

Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/

func main() {

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	shortestLength := findLadders(beginWord, endWord, wordList)
	fmt.Println(shortestLength)

	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	shortestLength = findLadders(beginWord, endWord, wordList)
	fmt.Println(shortestLength)

}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

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
	queue := [][]string{[]string{beginWord}}
	solution := [][]string{}
	for {
		var (
			queue2 [][]string
		)

		// If we reach the end of the queue with w/o returning, then transformations to
		// endWord does not exist
		if len(queue) == 0 {
			solution = [][]string{}
			break
		}

		for _, words := range queue {
			for idx, chars := range validTransf {
				for char := range chars {
					var (
						newWord string
					)

					// Need to replace the char at index idx with char
					// Recall that word[idx+1:] if idx == len(word) will be an empty string ""
					newWord = words[len(words)-1][:idx] + string(char) + words[len(words)-1][idx+1:]

					if newWord == endWord {
						// copy slice and append new word and add to solution
						// solution
						var soln = make([]string, len(words))
						copy(soln, words)
						soln = append(soln, newWord)
						solution = append(solution, soln)
					}

					// If created word is valid, then we need to copy the current sequence of words thus far
					// and check if the new word added to the sequence of words thus far is unique,
					// If not, then we skip/terminate the BFS there
					// If yes, then we continue the BFS by adding it to queueu2
					if _, ok := validWords[newWord]; ok {
						var addToQueue = make([]string, len(words))
						copy(addToQueue, words)
						addToQueue = append(addToQueue, newWord)

						if ok = allUniqueWords(addToQueue); ok {
							queue2 = append(queue2, addToQueue)
						}
					}
				}
			}
		}

		if len(solution) > 0 {
			return solution
		}

		queue = queue2
	}

	return solution
}

func allUniqueWords(words []string) bool {
	var (
		counter = make(map[string]int)
	)

	for _, word := range words {

		if _, ok := counter[word]; ok {
			return false
		}

		counter[word]++
	}

	return true
}
