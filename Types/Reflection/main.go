package main

import (
	"GoPractice/utils"
	"fmt"
	"reflect"
)

func main() {
	p := utils.GetProcess()
	ptr := reflect.ValueOf(p.Private)
	pc := ptr.Elem().FieldByName("Pc")
	priority := pc.FieldByName("Priority").Float()
	fmt.Println(float32(priority))
}
