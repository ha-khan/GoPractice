package utils

// BinaryTree ..
type BinaryTree struct {
	Root *TreeNode
}

// TreeNode ...
type TreeNode struct {
	Value       interface{}
	Right, Left *TreeNode
}

// NewBinaryTree is the constructor that returns
func NewBinaryTree() *BinaryTree {
	return nil
}
