package main

import "fmt"

func removeDuplicates(S string) string {
	if S == "" {
		return ""
	}
	stack := []uint8{}

	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, S[i])
	}

	return string(stack)
}

func main() {
	a := "aababaab"
	fmt.Println(removeDuplicates(a))
	b := "bbacbdd"
	fmt.Println(removeDuplicates(b))
	c := "aaaa"
	fmt.Println(removeDuplicates(c))
}
