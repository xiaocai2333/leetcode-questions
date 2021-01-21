// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

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
