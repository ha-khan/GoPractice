package main

import "fmt"

/*

Given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a
single digit. Add the two numbers and return the sum as a linked list.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

*/
// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		val   int
		carry int
		total *ListNode
	)
	val = l1.Val + l2.Val
	if val > 9 {
		val = val - 10
		carry = 1
	}

	total = &ListNode{Val: val, Next: nil}
	helper(l1.Next, l2.Next, total, carry)
	return total
}

func helper(l1, l2, total *ListNode, carry int) {

	var (
		val       int
		carryNext int
	)

	// base case
	if l1 == nil && l2 == nil && carry == 0 {
		total.Next = nil
	} else if l1 == nil && l2 == nil && carry > 0 {
		val = carry
		total.Next = &ListNode{Val: val, Next: nil}
		helper(nil, nil, total.Next, 0) // will hit base case
	} else if l1 != nil && l2 == nil {
		val = l1.Val + 0 + carry
		if val > 9 {
			val = val - 10
			carryNext = 1
		}
		total.Next = &ListNode{Val: val, Next: nil}
		helper(l1.Next, nil, total.Next, carryNext)
	} else if l1 == nil && l2 != nil {
		val = l2.Val + 0 + carry
		if val > 9 {
			val = val - 10
			carryNext = 1
		}
		total.Next = &ListNode{Val: val, Next: nil}
		helper(nil, l2.Next, total.Next, carryNext)
	} else {
		val = l1.Val + l2.Val + carry
		if val > 9 {
			val = val - 10
			carryNext = 1
		}
		total.Next = &ListNode{Val: val, Next: nil}
		helper(l1.Next, l2.Next, total.Next, carryNext)
	}

	return
}

func convertListNodeToString(total *ListNode) string {
	output := "["
	for t := total; t != nil; t = t.Next {
		if t.Next != nil {
			output += fmt.Sprintf("%+v,", t.Val)
		} else {
			output += fmt.Sprintf("%+v", t.Val)
		}
	}
	output += "]"

	return output
}

func buildListNodeFromSlice(l1 []int) *ListNode {
	head := &ListNode{Val: l1[0], Next: nil}
	buildListNodeFromSliceHelper(l1[1:], head)
	return head
}

func buildListNodeFromSliceHelper(l1 []int, total *ListNode) {
	if len(l1) == 0 {
		return
	}
	total.Next = &ListNode{Val: l1[0], Next: nil}
	buildListNodeFromSliceHelper(l1[1:], total.Next)
}

func main() {

	var (
		l1 *ListNode
		l2 *ListNode
	)

	l1 = buildListNodeFromSlice([]int{2, 4, 3})
	l2 = buildListNodeFromSlice([]int{5, 6, 4})
	fmt.Println(convertListNodeToString(l1) + " + " + convertListNodeToString(l2) + " = " + convertListNodeToString(addTwoNumbers(l1, l2)))

	l1 = buildListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = buildListNodeFromSlice([]int{9, 9, 9, 9})
	fmt.Println(convertListNodeToString(l1) + " + " + convertListNodeToString(l2) + " = " + convertListNodeToString(addTwoNumbers(l1, l2)))

	l1 = buildListNodeFromSlice([]int{0})
	l2 = buildListNodeFromSlice([]int{0})
	fmt.Println(convertListNodeToString(l1) + " + " + convertListNodeToString(l2) + " = " + convertListNodeToString(addTwoNumbers(l1, l2)))

	l1 = buildListNodeFromSlice([]int{9, 9, 9})
	l2 = buildListNodeFromSlice([]int{2})
	fmt.Println(convertListNodeToString(l1) + " + " + convertListNodeToString(l2) + " = " + convertListNodeToString(addTwoNumbers(l1, l2)))

}
