// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package main

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	parents []int
}

func newUnionFind2(n int) *UnionFind{
	parent := make([]int, n)
	for i, _ := range parent {
		parent[i] = i
	}

	return &UnionFind{
		parents: parent,
	}
}

func (uf *UnionFind) Find(x int) int{
	if uf.parents[x] != x {
		uf.parents[x] = uf.Find(uf.parents[x])
	}
	return uf.parents[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}

	uf.parents[rootY] = rootX

	return true
}

func abs2(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func dist2(x, y []int) int {
	return abs2(x[0]-y[0]) + abs2(x[1]-y[1])
}

func minCostConnectPoints2(points [][]int) int {
	var weights [][]int
	for i := 0; i < len(points); i++ {
		for j := i+1; j < len(points); j++ {
			weight := make([]int, 3)
			weight[0] = i
			weight[1] = j
			weight[2] = dist2(points[i], points[j])
			weights = append(weights, weight)
		}
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i][2] < weights[j][2]
	})
	result := 0
	uf := newUnionFind2(len(points))
	left := len(points)-1
	for _, weight := range weights {
		if uf.Union(weight[0], weight[1]) {
			left--
			result += weight[2]
		}
		if left <= 0 {
			break
		}
	}

	return result
}

func main() {
	points := [][]int{{3,12},{-2,5},{-4,1}}
	fmt.Println(minCostConnectPoints2(points))
}