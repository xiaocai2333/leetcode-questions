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