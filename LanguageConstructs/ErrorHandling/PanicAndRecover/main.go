package main

import (
	"fmt"
)

func main() {
	defer func() {
		if val := recover(); val != nil {
			switch t := val.(type) {
			case error:
				fmt.Printf("error.Error(): %s\n", t.Error())
			case string:
				fmt.Println(t)
			default:
				fmt.Println("Unknown type raised from panic!")
			}
		}
	}()

	// will panic since an array is not a slice
	// sort.Slice([5]int{1, 2, 3, 4, 5}, func(i, j int) bool {
	// 	return true
	// })
	//Compute()
}

func Compute() {
	panic("error")
}
