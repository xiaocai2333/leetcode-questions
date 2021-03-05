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

type MyQueue struct {
	in []int
	out []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.in = append(this.in, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		for i := len(this.in)-1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
			this.in = this.in[:i]
		}
	}
	popNum := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return popNum
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		return this.in[0]
	}

	return this.out[len(this.out)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.out) == 0 && len(this.in) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	a := []int{1,2,3}
	fmt.Println(a[0:0] == nil)
}
