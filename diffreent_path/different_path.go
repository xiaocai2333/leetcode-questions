package main

import "fmt"

func uniquePaths(m int, n int) int {
	var pathsNum [][]int

	for i := 0; i < m; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			if i==0 || j==0 {
				tmp = append(tmp, 1)
			}else {
				num := tmp[j-1]+pathsNum[i-1][j]
				tmp = append(tmp, num)
			}
		}
		pathsNum = append(pathsNum, tmp)
	}

	return pathsNum[m-1][n-1]
}

func main()  {
	m := 3
	n := 7

	r := uniquePaths(m, n)
	fmt.Println(r)
	return
}