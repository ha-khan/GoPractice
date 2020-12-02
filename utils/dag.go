package utils

type DAG struct {
	nodes []*DagNode
}

type DagNode struct {
	Value   interface{}
	OutEdge []*DagNode
}
