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

func addToArrayForm(A []int, K int) []int {
	n := len(A)
	var temp []int
	j := n-1
	for {
		if j >= 0{
			K = K + A[j]
		}
		temp = append(temp, K % 10)
		K = K / 10
		j--
		if K / 10 == 0 && K % 10 == 0 && j < 0 {
			break
		}
	}
	m := len(temp)
	result := make([]int, m)
	for i, value := range temp {
		result[m-1-i] = value
	}
	return result
}

func main() {
	test1 := []int{9, 9, 9, 9, 9}
	k := 1
	fmt.Println(addToArrayForm(test1, k))

	test2 := []int{1, 2, 0, 0}
	k2 := 34
	fmt.Println(addToArrayForm(test2, k2))
}