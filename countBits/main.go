package main

import "fmt"

func countBits(num int) []int {
	bits := make([]int, num+1)
	bits[0] = 0
	for i := 1; i < num+1; i++ {
		bits[i] = bits[i>>1] + i&1
	}

	return bits
}

func main() {
	fmt.Println(countBits(5))
	fmt.Println(countBits(0))
	fmt.Println(countBits(10))
}
