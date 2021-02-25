package main

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	candyNum := 0
	candyForEachOne := make([]int, n)
	candyForEachOne[0] = 1
	for i := 1; i < n; i++{
		if ratings[i] > ratings[i-1] {
			candyForEachOne[i] = candyForEachOne[i-1]+1
		} else {
			candyForEachOne[i] = 1
			if ratings[i] < ratings[i-1] && candyForEachOne[i] >= candyForEachOne[i-1] {
				for j := i; j > 0; j-- {
					if ratings[j] < ratings[j-1] && candyForEachOne[j] >= candyForEachOne[j-1] {
						candyForEachOne[j-1]++
					}else {
						break
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		candyNum += candyForEachOne[i]
	}

	return candyNum
}
