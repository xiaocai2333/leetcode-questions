package nextBiggerNum

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	stack := []int{}
	for i := 0; i < 2*n-1; i++ {
		for len(stack) >0 && nums[stack[len(stack)-1]] < nums[i%n] {
			result[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return result
}
