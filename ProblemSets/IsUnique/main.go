package main

import (
	"fmt"
)

/*

Assume it is only composed of ASCII characters

*/
func isUnique(str string) bool {
	for idx, val := range str {
		for idx2, val2 := range str {
			if idx != idx2 && val == val2 {
				return false
			}
		}
	}

	return true
}

func isUniqueMemory(str string) bool {
	cache := make(map[string]interface{})

	for _, val := range str {
		if _, ok := cache[string(val)]; ok {
			return false
		}
		cache[string(val)] = true
	}

	return true
}

func main() {
	fmt.Println(isUnique("strings"))
	fmt.Println(isUniqueMemory("string"))
}
