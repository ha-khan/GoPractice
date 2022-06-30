package main

import "GoPractice/utils"

func main() {

	List := utils.NewLinkedList(123)
	List.PushBack(456)
	List.PushBack(789)
	List.PushBack(1011)

	List.PrintList()

	return
}
