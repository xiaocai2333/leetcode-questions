package main

import "fmt"

func isMonotonic(A []int) bool {
	inc, dec := false, false
	for i := 0; i < len(A)-1; i++ {
		if A[i+1] - A[i] > 0 {
			inc = true
		}
		if A[i+1] - A[i] < 0 {
			dec = true
		}
		if inc && dec {
			return false
		}
	}

	return true
}

func main() {
	a := []int{1,3,3,4,5}
	b := []int{6,4,4,5,3,2}
	c := []int{1,2,3,3,4,4,4,4,4,4,4,3,-1,-1,-1}
	fmt.Println(isMonotonic(a))
	fmt.Println(isMonotonic(b))
	fmt.Println(isMonotonic(c))
}
