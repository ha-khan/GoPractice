package main

func main() {}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return univalTreeHelper(root, root.Val)
}

func univalTreeHelper(root *TreeNode, val int) bool {

	var (
		left  = true
		right = true
	)

	if root == nil {
		return true
	}

	if root.Val != val {
		return false
	}

	left = left && univalTreeHelper(root.Left, val)
	right = right && univalTreeHelper(root.Right, val)

	return left && right
}
