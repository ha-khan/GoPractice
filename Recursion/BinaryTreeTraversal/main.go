package main

// TreeNode ...
type TreeNode struct {
	Left, Right *TreeNode
	Data        interface{}
}

func main() {
	var tree TreeNode

	tree.Data = "100"

	tree.Left = &TreeNode{Left: nil, Right: &TreeNode{Data: "200"}, Data: "300"}

	tree.Right = &TreeNode{Data: "400"}
}
