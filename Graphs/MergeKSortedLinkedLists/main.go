package main

import (
	"GoPractice/utils"
)

func main() {

	L1 := utils.NewLinkedList(1)
	L1.PushBack(4)
	L1.PushBack(5)

	L2 := utils.NewLinkedList(1)
	L2.PushBack(3)
	L2.PushBack(4)

	L3 := utils.NewLinkedList(2)
	L3.PushBack(6)

	sorted := mergeKLists([]*utils.LinkedList{L1, L2, L3})

	sorted.PrintList()
}

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []
*/
func mergeKLists(lists []*utils.LinkedList) *utils.LinkedList {

	var (
		sorted          *utils.LinkedList
		smallestNode    *utils.ListNode
		smallestNodeIdx int
		listNodes       = make([]*utils.ListNode, 0)
		nilNodes        int
	)

	getIntVal := func(val interface{}) int {
		intVal, _ := val.(int)
		return intVal
	}

	// First populate []*utils.ListNode, LinkedList is a wrapper object
	for _, node := range lists {
		listNodes = append(listNodes, node.Root)
	}

	// TODO: Clean this up a bit
	for nilNodes != len(listNodes) {
		nilNodes = 0
		for idx, thisNode := range listNodes {

			// We've reached the end of this LinkedList so skip remainder of block
			if thisNode == nil {
				nilNodes++
				continue
			}

			if smallestNode == nil {
				smallestNode = thisNode
				smallestNodeIdx = idx
				nilNodes++
			} else {

				smallestNodeValue := getIntVal(smallestNode.Value)
				thisNodeValue := getIntVal(thisNode.Value)

				// Keep track of smallest seen val of a given thisNode stored in smallestNode
				if smallestNodeValue > thisNodeValue {
					smallestNode = thisNode
					smallestNodeIdx = idx
				}
			}
		}

		if nilNodes == len(listNodes) {
			break
		}

		if sorted == nil {
			sorted = utils.NewLinkedList(smallestNode.Value)
		} else {
			sorted.PushBack(smallestNode.Value)
		}

		// Mutate
		listNodes[smallestNodeIdx] = listNodes[smallestNodeIdx].Next
		smallestNode = nil
	}

	return sorted
}
