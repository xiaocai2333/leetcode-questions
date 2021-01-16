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
)

func hitBricks(grid [][]int, hits [][]int) []int {
	rows, cols := len(grid), len(grid[0])
	var copyGrid [][]int
	parents := make([]int, rows*cols+1)
	size := make([]int, rows*cols+1)
	results := make([]int, len(hits))
	for i  := 0; i < len(grid); i++{
		var row []int
		row = append(row, grid[i][:]...)
		copyGrid = append(copyGrid, row)
	}

	for i := range parents {
		parents[i] = i
		size[i] = 1
	}
	var getIndex func(int, int) int
	getIndex = func(x, y int) int {
		return x*cols+y
	}

	var find func(int)int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}

		return parents[x]
	}

	var union func(int, int)
	union = func(x int, y int) {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			parents[rootX] = rootY
			size[rootY] += size[rootX]
		}
	}

	var findSize func(int)int
	findSize = func(x int) int {
		rootX := find(x)
		return size[rootX]
	}

	gridRoot := rows*cols

	for i := 0; i < len(copyGrid[0]); i++ {
		if copyGrid[0][i] == 1 {
			union(getIndex(0, i), gridRoot)
		}
	}

	for i := 0; i < len(hits); i++ {
		copyGrid[hits[i][0]][hits[i][1]] = 0
	}

	for i := 1; i < len(copyGrid); i++ {
		for j := 0; j < len(copyGrid[i]); j++ {
			if copyGrid[i][j] == 1 {
				if copyGrid[i-1][j] == 1 {
					union(getIndex(i-1, j), getIndex(i, j))
				}

				if j > 0 && copyGrid[i][j-1] == 1 {
					union(getIndex(i, j-1), getIndex(i, j))
				}
			}
		}
	}

	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]

		if grid[x][y] == 0 {
			continue
		}
		if x == 0 {
			union(getIndex(x, y), gridRoot)
		}
		originSize := findSize(getIndex(x, y))

		if y > 0 && copyGrid[x][y-1] == 1 {
			union(getIndex(x, y-1), getIndex(x, y))
		}
		if x > 0 && copyGrid[x-1][y] == 1 {
			union(getIndex(x-1, y), getIndex(x, y))
		}
		if y+1 < cols && copyGrid[x][y+1] == 1 {
			union(getIndex(x, y+1), getIndex(x, y))
		}
		if x+1 < rows && copyGrid[x+1][y] == 1 {
			union(getIndex(x+1, y), getIndex(x, y))
		}

		currentSize := findSize(getIndex(x,y))

		if currentSize-originSize-1 < 0 {
			results[i] = 0
		}else {
			results[i] = currentSize-originSize-1
		}

		copyGrid[x][y] = 1
	}

	return results
}

func main () {
	grid1 := [][]int{{1,0,0,0}, {1,1,0,0}}
	hits1 := [][]int{{1,1}, {1,0}}

	result1 := hitBricks(grid1, hits1)
	//expacl1 := []int{0,0}
	fmt.Println(result1)

	grid2 := [][]int{{1,0,0,0}, {1,1,1,0}}
	hits2 := [][]int{{1,0}}

	result2 := hitBricks(grid2, hits2)
	//expacl2 := []int{2}
	fmt.Println(result2)
}