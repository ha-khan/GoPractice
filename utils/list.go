package utils

import "fmt"

// LinkedList is the wrapper class for recursively defined ListNodes classes
type LinkedList struct {
	Root *ListNode
}

// ListNode is a composite type that
type ListNode struct {
	Value any
	Next  *ListNode
}

// ListIterator
type ListIterator struct {
	curr *ListNode
}

// NewLinkedList is the default constructor for creating a new LinkedList instance
// Where the root is pointing to a ListNode that wraps around the inputted val interface{}
func NewLinkedList(val any) *LinkedList {
	return &LinkedList{Root: &ListNode{Value: val, Next: nil}}
}

// PushBack will insert the inputted item at the end of the LinkedList
// Time Complexity ~ O(n)
func (l *LinkedList) PushBack(val any) {
	var (
		cursor *ListNode = l.Root
	)

	for cursor.Next != nil {
		cursor = cursor.Next
	}

	// Just reassigns this pointer variable to the one that &ListNode
	// cursor = &ListNode{Value: val, Next: nil}

	cursor.Next = new(ListNode)
	cursor.Next.Value = val
	cursor.Next.Next = nil
}

// PrintList ...
func (l LinkedList) PrintList() {
	for curr := l.Root; curr != nil; curr = curr.Next {
		fmt.Print(fmt.Sprintf("%d->", curr.Value))
	}

	fmt.Println("NULL")
}

func (l LinkedList) Begin() *ListIterator {
	return &ListIterator{
		curr: l.Root,
	}
}
