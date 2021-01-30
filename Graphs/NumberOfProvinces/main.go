package main

import (
	"fmt"
)

func main() {

	var (
		data [][]int
	)

	data = [][]int{[]int{1, 1, 0}, []int{1, 1, 0}, []int{0, 0, 1}}
	fmt.Printf("Matrix: %+v, Provinces: %d\n", data, findCircleNum(data))

	data = [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}}
	fmt.Printf("Matrix: %+v, Provinces: %d\n", data, findCircleNum(data))

	data = [][]int{[]int{1, 1, 1}, []int{0, 0, 0}, []int{0, 0, 0}}
	fmt.Printf("Matrix: %+v, Provinces: %d\n", data, findCircleNum(data))

	data = [][]int{[]int{1, 1, 0, 0, 0, 0}, []int{1, 1, 0, 0, 0, 0}, []int{0, 0, 1, 1, 1, 0},
		[]int{0, 0, 1, 1, 0, 0}, []int{0, 0, 1, 0, 1, 0}, []int{0, 0, 0, 0, 0, 1}}
	fmt.Printf("Matrix: %+v, Provinces: %d\n", data, findCircleNum(data))s
}

/*

There are n cities. Some of them are connected, while some are not.

If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.


Example 1:
Input: isConnected = [[1,1,0],
					  [1,1,0],
					  [0,0,1]]
Output: 2


Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

*/
func findCircleNum(isConnected [][]int) int {

	var (
		provinces int
		visited   = make(map[int]struct{}) // set to keep track of indices 0 -> len(isConnected) - 1
	)

	for _, row := range isConnected {
		for idy, colVal := range row {

			// check if value at location isConnected[idx][idy] == 1
			if colVal != 1 {
				continue
			}

			// check to see if node has been visited yet, if not then start DFS procedure
			// DFS returns when it can no longer find newly connected nodes/towns
			if _, ok := visited[idy]; !ok {
				dfs(idy, 0, isConnected, visited)
				provinces++
			}

		}
	}

	return provinces
}

func dfs(i, j int, isConnected [][]int, visited map[int]struct{}) {

	// mark i as visited
	visited[i] = struct{}{}

	for idy, val := range isConnected[i] {
		if val != 1 {
			continue
		}
		if _, ok := visited[idy]; !ok {
			dfs(idy, 0, isConnected, visited)
		}
	}
}
