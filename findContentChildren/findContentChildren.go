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

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) ==0 || len(g) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	count := 0
	for i != len(g) && j != len(s) {
		if s[j] >= g[i] {
			count++
			i++
			j++
		}else {
			j++
		}
	}
	return count
}
