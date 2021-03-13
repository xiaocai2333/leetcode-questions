package main

import "fmt"

func isValidSerialization(preorder string) bool {
	leaf := 0
	root := 0
	i := 0
	for ; i < len(preorder);  {
		switch preorder[i] {
		case ',':
			i++
		case '#':
			leaf++
			i++
		default:
			for i < len(preorder) {
				if preorder[i] >= '0' && preorder[i] <= '9' {
					i++
				}else {
					root++
					i++
					break
				}
			}
		}
		if leaf > root + 1 {
			return false
		}
		if leaf == root + 1 && i < len(preorder)-1{
			return false
		}
	}
	if leaf	!= root + 1 {
		return false
	}
	return true
}

func main()  {
	s := "19,223,4,#,#,1,#,#,3332,#,1116,#,#"
	fmt.Println(isValidSerialization(s))
	s2 := "#"
	fmt.Println(isValidSerialization(s2))
	s3 := "9,#,#,1"
	fmt.Println(isValidSerialization(s3))
}
