package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	if n <= 0 {
		return nil
	}

	m := len(matrix[0])
	var result [][]int
	for i := 0; i < m; i++ {
		var col []int
		for j := 0; j < n; j++ {
			col = append(col, matrix[j][i])
		}
		result = append(result, col)
	}

	return result
}

func main() {
	matrix := [][]int{{1,2,3},{4,5,6}}
	fmt.Println(transpose(matrix))
}
