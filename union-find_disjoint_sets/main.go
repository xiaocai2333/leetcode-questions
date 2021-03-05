package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {

	//var parent map[int]int
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(int)int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	var union func(int, int) bool
	union = func(from int, to int) bool {
		if find(from) == find(to) {
			return false
		}

		parent[find(to)] = parent[find(from)]
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}

	return nil
}

func main() {
	edges := [][]int{{1,2},{2,3},{3,4},{1,4},{1,5}}
	edges2 := [][]int{{1,4}, {3, 4}, {1,3}, {1,2}, {4,5}}

	fmt.Println(findRedundantConnection(edges))
	fmt.Println(findRedundantConnection(edges2))
}

