package main

import (
	"fmt"
	"math/bits"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	puzzleLength := 7
	cnt := map[int]int{}
	n := len(puzzles)
	result := make([]int, n)
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(uint(mask)) <= puzzleLength {
			cnt[mask]++
		}
	}
	for i, puzzle := range puzzles {
		puzzleFirst := 1 << (puzzle[0] - 'a')
		mask := 0

		for _, ch := range puzzle[1:] {
			mask |= 1 << (ch - 'a')
		}

		temp := mask
		for {
			result[i] += cnt[temp|puzzleFirst]
			temp = (temp - 1) & mask
			if temp == mask {
				break
			}
		}
	}
	return result
}

func main() {
	words := []string{"aaaa","asas","able","ability","actt","actor","access"}
	puzzles := []string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}
	//result := [1,1,3,2,4,0]

	fmt.Println(findNumOfValidWords(words, puzzles))
}
