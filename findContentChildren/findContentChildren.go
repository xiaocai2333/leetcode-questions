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
