package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	return max(f...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	envelopes := [][]int{{1,2},{3,7},{4,5},{5,6}}
	fmt.Println(maxEnvelopes(envelopes))
}
