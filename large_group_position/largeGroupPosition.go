package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	index := 0
	count := 1
	var result [][]int
	n := len(s)
	for i := 1; i < n; i++{
		if s[i] == s[i-1] {
			count++
		} else {
			if count >= 3 {
				result = append(result, []int{index, i-1})
			}

			index = i
			count = 1
		}
		if i == n-1 {
			if count >= 3 {
				result = append(result, []int{index, i})
			}
		}
	}

	return result
}

func main() {
	s := "abcdddeeeeaabbbcd"
	//输出：[[3,5],[6,9],[12,14]]
	fmt.Println(largeGroupPositions(s))
}