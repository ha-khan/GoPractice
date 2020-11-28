package utils

// BinaryTreeNode is node
type BinaryTreeNode struct {
	Val         interface{}
	Left, Right *BinaryTreeNode
}

// BinaryTree is a wrapper struct around the BinaryTreeNode struct and contains a single
type BinaryTree struct {
	Root *BinaryTreeNode
}
