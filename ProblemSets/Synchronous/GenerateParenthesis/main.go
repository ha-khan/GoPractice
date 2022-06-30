package main

import (
	"fmt"
)

type GenerateParenthesis int32

func (g GenerateParenthesis) Generate() string {

	func() {}()
	return ""
}

func main() {

	var g GenerateParenthesis

	g = 213

	y := g.Generate

	y()

	fmt.Println(g)

}
