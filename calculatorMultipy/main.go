package main

import "fmt"

func calculate(s string) (ans int) {
	var stack []int
	preOps := 1
	i := 0
	for ; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			preOps = 1
			i++
		case '-':
			preOps = -1
			i++
		case '*':
			i++
			num := 0
			for ; i < len(s); i++ {
				if s[i] == ' ' {
					continue
				}
				if s[i] >= '0' && s[i] <= '9' || s[i] == ' ' {
					num = num*10 + int(s[i]-'0')
				}else {
					break
				}
			}
			stack[len(stack)-1] = stack[len(stack)-1] * num
		case '/':
			i++
			num := 0
			for ; i < len(s); i++ {
				if s[i] == ' ' {
					continue
				}
				if s[i] >= '0' && s[i] <= '9' || s[i] == ' ' {
					num = num*10 + int(s[i]-'0')
				}else {
					break
				}
			}
			stack[len(stack)-1] = stack[len(stack)-1] / num
			preOps = 1
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			stack = append(stack, num*preOps)
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}

	return result
}

func main()  {
	s := "1 + 200 * 3 / 5 - 7 * 8 + 11"
	fmt.Println(calculate(s))
	s2 := "3+2*2"
	fmt.Println(calculate(s2))
	s3 := "0-1"
	fmt.Println(calculate(s3))
}
