package main

import (
	"fmt"
)

/*
Three types of edits that can be performed on strings:
	1. Insert Char
	2. Remove Char
	3. Replace Char


	Give two strings, write a function to check if they are	one
	edit away or zero edits away.

	pale,  ple   -> true
	pales, pale  -> true
	pale,  bale  -> true
	pale,  bake  -> false
*/

func checkOneEdit(a, b string) bool {

	var edits int

	for i, j := 0, 0; i < len(a) && j < len(b) && edits <= 1; {
		if a[i] != b[j] {
			// If chars are not the same, check which is greater
			if len(a) != len(b) {
				edits++
				if len(a) > len(b) {
					i++
				} else {
					j++
				}
			} else {
				edits++
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}

	return edits <= 1

}

func main() {

	fmt.Println(checkOneEdit("pale", "ple"))
	fmt.Println(checkOneEdit("pales", "pale"))
	fmt.Println(checkOneEdit("pale", "bale"))
	fmt.Println(checkOneEdit("pale", "bake"))

}
