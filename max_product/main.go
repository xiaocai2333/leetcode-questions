package main

import "fmt"

func sort(a,b,c int) (int, int, int) {
	if a >= b {
		if b >= c {
			return a, b, c
		}
		if a >= c {
			return a, c, b
		}
		return c, a, b
	}
	if a >= c {
		return b, a, c
	}
	if c >= b {
		return c, b, a
	}

	return b, c, a
}

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var biggest, secondLargest, thirdLargest, smallest, secondSmallest int

	biggest, secondLargest, thirdLargest = sort(nums[0], nums[1], nums[2])
	smallest = thirdLargest
	secondSmallest = secondLargest

	for _, num := range nums[3:] {
		if num > biggest {
			thirdLargest = secondLargest
			secondLargest = biggest
			biggest = num
		}else if num > secondLargest {
			thirdLargest = secondLargest
			secondLargest = num
		}else if num > thirdLargest {
			thirdLargest = num
		}

		if num < smallest {
			secondSmallest = smallest
			smallest = num
		} else if num < secondSmallest {
			secondSmallest = num
		}
	}

	positiveProduct := biggest*secondLargest*thirdLargest
	negativeProduct := biggest*smallest*secondSmallest
	if positiveProduct > negativeProduct {
		return positiveProduct
	}

	return negativeProduct
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(maximumProduct(nums))
}
