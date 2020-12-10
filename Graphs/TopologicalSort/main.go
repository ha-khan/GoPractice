package main

import (
	"GoPractice/utils"
)

func main() {

	v3 := utils.NewDagNode(utils.OutEdges{}, 3)
	v5 := utils.NewDagNode(utils.OutEdges{}, 5)
	v4 := utils.NewDagNode(utils.OutEdges{v5}, 4)
	v2 := utils.NewDagNode(utils.OutEdges{v4, v3}, 2)
	v6 := utils.NewDagNode(utils.OutEdges{v4}, 6)
	v1 := utils.NewDagNode(utils.OutEdges{v2, v4}, 1)
	v0 := utils.NewDagNode(utils.OutEdges{v1}, 0)

	DAG := utils.NewDAG()

	DAG.InsertNode(v3)
	DAG.InsertNode(v5)
	DAG.InsertNode(v6)
	DAG.InsertNode(v0)
	DAG.InsertNode(v2)
	DAG.InsertNode(v4)
	DAG.InsertNode(v1)

	// v0 --> v1 --> v2 --> v3
	//          \  /
	//           v4 --> v5
	//          /
	//         v6

	// v0 v1 v2 v6 v4 v3 v5

	DAG.DFS()

}
