package main

import "fmt"

func rotate(matrix [][]int)  {
	var tempMatrix [][]int
	for i := 0; i < len(matrix[0]); i++ {
		var temp []int
		for j := 0; j < len(matrix); j++{
			temp = append(temp, matrix[len(matrix)-1-j][i])
		}
		tempMatrix = append(tempMatrix, temp)
	}
	copy(matrix, tempMatrix)
}


func main(){
	var test1 = [][]int{{1,2,3},{4,5,6},{7,8,9}}

	rotate(test1)
	fmt.Println(test1)
}