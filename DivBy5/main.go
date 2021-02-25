package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	currDecimalNum := 0
	result := make([]bool, len(A))

	for i, v := range A {
		currDecimalNum = (currDecimalNum << 1 + v) % 5
		if currDecimalNum == 0 {
			result[i] = true
			continue
		}
		result[i] = false
	}

	return result
}

func main() {
	A := []int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0}
	fmt.Println(prefixesDivBy5(A))
}
