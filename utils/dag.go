package utils

import (
	"fmt"
)

// NewDAG returns a ptr to a DAG object
func NewDAG() *DAG {
	return new(DAG)
}

// NewDagNode returns a NewDagNode
func NewDagNode(dependencies []*DagNode, val interface{}) (node *DagNode) {
	node = new(DagNode)
	node.Value = val
	node.OutEdges = append(node.OutEdges, dependencies...)
	return
}

// OutEdges ...
type OutEdges []*DagNode

// DAG ... is the class definition for DAG objects
type DAG struct {
	nodes []*DagNode
}

// DagNode is the composite
type DagNode struct {
	Value interface{}
	OutEdges
}

// DFS ...
func (d DAG) DFS() {

	var (
		// keeps track of nodes visited thus far for DFS
		cache = make(map[*DagNode]struct{})
	)

	for _, node := range d.nodes {
		if _, ok := cache[node]; !ok {
			d.dfsHelper(node, cache)
		}
	}
}

// BFS ...
func (d DAG) BFS() {

}

func (d DAG) dfsHelper(node *DagNode, cache map[*DagNode]struct{}) {

	if _, ok := cache[node]; ok {
		return
	}

	// Mark node as visited
	cache[node] = struct{}{}
	for _, nodeout := range node.OutEdges {
		d.dfsHelper(nodeout, cache)
	}
	fmt.Print(fmt.Sprintf("v%d<-", node.Value))
}

// InsertNode inserts a node into the DAG,
func (d *DAG) InsertNode(node *DagNode) {
	d.nodes = append(d.nodes, node)
}
