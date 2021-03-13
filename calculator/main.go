package main

import (
	"fmt"
)

func calculate(s string) int {
	result := 0
	n := len(s)
	ops := []int{1}
	sign := 1
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result += sign*num
		}
	}

	return result
}

func main() {
	s := "(1+(4+5+2)-3)+(6+8)"
	a := "1 + 1"
	b := "-1"
	c := "- (3 + (4 + 5))"
	fmt.Println(calculate(c))
	fmt.Println(calculate(b))
	fmt.Println(calculate(a))
	fmt.Println(calculate(s))
}
