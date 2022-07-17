package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5}
	c := []int{6}

	lt := [][]int{a, b, c}

	solution := CartesianProduct(lt...)

	fmt.Println(solution)
}

func CartesianProduct(input ...[]int) [][]int {
	fn := func(a []int, b [][]int) [][]int {
		var sol [][]int
		for _, v1 := range a {
			for _, v2 := range b {
				temp := make([]int, len(v2))
				copy(temp, v2)
				temp = append(temp, v1)
				sol = append(sol, temp)
			}
		}

		return sol
	}

	var init = [][]int{{}}
	for _, v := range input {
		init = fn(v, init)
	}

	return init
}
