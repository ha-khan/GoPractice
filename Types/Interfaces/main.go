package main

import "fmt"

type Query struct {
	method string
	data   []byte
}

func (q *Query) DoGet() error {
	return nil
}

func main() {

	var query *Query

	// Allows you to quickly check whether a passed in concrete type implements a method
	func(q interface{}) {
		if _, ok := q.(interface {
			DoGet() error
		}); ok {
			fmt.Println("input q implements DoGet()")
		}
	}(query)

}
