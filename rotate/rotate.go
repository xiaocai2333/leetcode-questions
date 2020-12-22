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

func rotate(matrix [][]int)  {
	var tempMatrix [][]int
	for i := 0; i < len(matrix[0]); i++ {
		var temp []int
		for j := 0; j < len(matrix); j++{
			temp = append(temp, matrix[len(matrix)-1-j][i])
		}
		tempMatrix = append(tempMatrix, temp)
	}
	copy(matrix, tempMatrix)
}


func main(){
	var test1 = [][]int{{1,2,3},{4,5,6},{7,8,9}}

	rotate(test1)
	fmt.Println(test1)
}