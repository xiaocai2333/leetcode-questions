package main

import "fmt"

type NumArray struct {
	sums []int
}


func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	return NumArray{
		sums: sums,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.sums[j+1] - this.sums[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main()  {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)
	fmt.Println(numArray.sums)
	fmt.Println(numArray.SumRange(0,2))
	fmt.Println(numArray.SumRange(2,5))
	fmt.Println(numArray.SumRange(0,5))
}