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

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	candyNum := 0
	candyForEachOne := make([]int, n)
	candyForEachOne[0] = 1
	for i := 1; i < n; i++{
		if ratings[i] > ratings[i-1] {
			candyForEachOne[i] = candyForEachOne[i-1]+1
		} else {
			candyForEachOne[i] = 1
			if ratings[i] < ratings[i-1] && candyForEachOne[i] >= candyForEachOne[i-1] {
				for j := i; j > 0; j-- {
					if ratings[j] < ratings[j-1] && candyForEachOne[j] >= candyForEachOne[j-1] {
						candyForEachOne[j-1]++
					}else {
						break
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		candyNum += candyForEachOne[i]
	}

	return candyNum
}
