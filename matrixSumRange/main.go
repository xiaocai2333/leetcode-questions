package main

import (
	"fmt"
)

type NumMatrix struct {
	sums [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sums[i] = make([]int, len(matrix[i])+1)
	}
	for i, array := range matrix {
		for j, num := range array {
			sums[i][j+1] = sums[i][j] + num
		}
	}

	return NumMatrix{
		sums: sums,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for i := row1; i <= row2; i++ {
		result += this.sums[i][col2+1] - this.sums[i][col1]
	}
	return result
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5}}

	matrix2 := Constructor(matrix)

	fmt.Println(matrix2)
	fmt.Println(matrix2.SumRegion(2,1,4,3))
}

