package main

import (
	"fmt"
	"net"
)

//Variables of type 'interface' can hold 'values' of any type that implements all the methods defined by the 'interface'
type Query struct {
	method string
	data   []byte
}

func (q *Query) DoGet() error {
	return nil
}

// Nested anonymous struct definitions; similar to nested class definitions
type Config struct {
	Networking struct {
		IP   net.IP
		Port int32
	}
}

func main() {

	// Allows you to quickly check whether a passed in concrete type implements a method
	var query *Query
	func(q interface{}) {
		if _, ok := q.(interface {
			DoGet() error
		}); ok {
			fmt.Println("input q implements DoGet()")
		}
	}(query)

	c := Config{
		Networking: struct {
			IP   net.IP
			Port int32
		}{
			IP:   net.IP{8, 8, 8, 8},
			Port: 42300,
		},
	}
	fmt.Println(c.Networking.IP)
}
