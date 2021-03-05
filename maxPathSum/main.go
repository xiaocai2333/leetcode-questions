package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	var oneSideMax func(root *TreeNode) int
	oneSideMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := max(0, oneSideMax(root.Left))
		rightSum := max(0, oneSideMax(root.Right))

		result = max(result, leftSum+rightSum+root.Val)

		return max(leftSum, rightSum) + root.Val
	}

	rootSum := oneSideMax(root)
	return max(result, rootSum)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(maxPathSum(node))

	node2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: -1,
		},
	}

	fmt.Println(maxPathSum(node2))
}
