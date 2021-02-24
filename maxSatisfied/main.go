package maxSatisfied

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	var total int
	for i := 0; i < n; i++ {
		total -= customers[i] * (grumpy[i] - 1)
	}

	result := total
	tmp := total

	for i := 0; i < n-X+1; i++ {
		for j := i; j < i+X; j++ {
			if grumpy[j] == 1 {
				tmp += customers[j]
			}
		}
		if tmp > result {
			result = tmp
		}
		tmp = total
	}

	return result
}
