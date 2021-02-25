package main

import "fmt"

/**
* Definition for a binary tree node.
*/
type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
 }

func zigzagLevelOrder(root *TreeNode) [][]int {
	isEnd := false
	var result [][]int
	h := 1
	var roots []*TreeNode
	i := 0
	Loop:
		for {
			var preLayerVal []int
			if root == nil {
				isEnd = true
				break Loop
			}
			isEnd = true
			n := len(roots)
			if h == 1 {
				preLayerVal = append(preLayerVal, root.Val)
				roots = append(roots, root)
				isEnd = false
			} else {
				for j := n - 1; j >= i; j-- {
					root := roots[j]
					if h % 2 == 0 {
						if root.Right != nil {
							preLayerVal = append(preLayerVal, root.Right.Val)
							roots = append(roots, root.Right)
							isEnd = false
						}
						if root.Left != nil {
							preLayerVal = append(preLayerVal, root.Left.Val)
							roots = append(roots, root.Left)
							isEnd = false
						}
					} else {
						if root.Left != nil {
							preLayerVal = append(preLayerVal, root.Left.Val)
							roots = append(roots, root.Left)
							isEnd = false
						}
						if root.Right != nil {
							preLayerVal = append(preLayerVal, root.Right.Val)
							roots = append(roots, root.Right)
							isEnd = false
						}
					}
				}
			}
			if isEnd == true {
				break Loop
			}
			h++
			i = n
			result = append(result, preLayerVal)

		}
	return result
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 7,
				Left: nil,
				Right: nil,
			},
		},
	}

	fmt.Println(zigzagLevelOrder(root))
}

