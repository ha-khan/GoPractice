package main

import (
	"fmt"
	"testing"
)

type soln [2]string

func Test1(t *testing.T) {

	// a := soln{"hello", "world"}

	// str := [][2]string{a}

	str := longestPalindromicSubStringFast("abcdefghtacocatijklol")

	if str != "tacocat" {
		t.Error(fmt.Sprintf("Expected str: %s, but got %s", "tacocat", str))
	}

}
