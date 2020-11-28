package main

// TreeNode ...
type BinaryTreeNode struct {
	Left, Right *BinaryTreeNode
	Data        interface{}
}

func main() {
	var tree BinaryTreeNode

	tree.Data = "100"

	tree.Left = &BinaryTreeNode{Left: nil, Right: &BinaryTreeNode{Data: "200"}, Data: "300"}

	tree.Right = &BinaryTreeNode{Data: "400"}
}
