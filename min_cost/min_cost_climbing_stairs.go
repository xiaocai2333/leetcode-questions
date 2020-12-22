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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))

	cost2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(minCostClimbingStairs(cost2))
	return
}
