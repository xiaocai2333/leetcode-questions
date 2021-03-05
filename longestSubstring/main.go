package main

import (
	"fmt"
	"strings"
)

func longestSubstring(s string, k int) int {
	cnt := make([]int, 26)

	for _, c := range s {
		cnt[c-'a']++
	}

	var split uint8
	for c, v := range cnt[:] {
		if 0 < v && v < k {
			split = 'a'+uint8(c)
		}
	}

	if split == 0 {
		return len(s)
	}
	result := 0
	for _, subStr := range strings.Split(s, string(split)) {
		result = max(result, longestSubstring(subStr, k))
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "ababbc"
	k := 2

	fmt.Println(longestSubstring(s, k))
}
