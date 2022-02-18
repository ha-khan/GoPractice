package main

import (
	"GoPractice/utils"
	"fmt"
	"reflect"
)

/*

https://go.dev/blog/laws-of-reflection



At the basic level, reflection (in go) is just a mechanism to examine the type and value pair stored inside an interface variable.

the laws of reflection:

    Reflection goes from interface value to reflection object.

    Reflection goes from reflection object to interface value.

    To modify a reflection object, the value must be settable.

*/
func main() {
	// A fairly nested structure where the internals are not exposed to importing packages
	p := utils.GetProcess()

	// Luckily the structure is written in such a way that allows us to use reflection to "traverse" down the structure
	ptr := reflect.ValueOf(p.Private)

	// The returned concrete type is a ptr to another structure which has its type definition in the utils package
	// and thus not exposed to this package. however we can peek into those files since we're importing that package
	// and we can extract specific fields by ref their names
	pc := ptr.Elem().FieldByName("Pc")

	// pc is another struct and we want to do the same thing and extract a specific attribute
	// which is a primitive type, which we typecast to float64 (which we know by reading the imported package source)
	priority := pc.FieldByName("Priority").Float()

	fmt.Println(float32(priority))
}
