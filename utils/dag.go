package utils

import (
	"fmt"
)

// Node is a vertex in a DAG which has references to other Nodes (Out going edges)
type Node struct {
	Val     int
	OutEdge []*Node
}

// DAG is a list of Nodes; assumption is that the user will construct is accordingly
type DAG struct {
	Nodes []Node
}

// NewDAG returns a new DAG instance
func NewDAG() DAG {

	return nil
}

// PrintOrder prints the order of the nodes in the DAG starting at index 0
func (d *DAG) PrintOrder() {

	var visited = make(map[*Node]interface{})

	for _, node := range *d {
		visited[&node] = true
		fmt.Println(node.Val)
		for _, out := range node.OutEdge {
			fmt.Println(out.Val)

		}
	}
}
