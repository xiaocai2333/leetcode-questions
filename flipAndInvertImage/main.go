package flipAndInvertImage

func flipAndInvertImage(A [][]int) [][]int {
	n := len(A)
	if n == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		for j := 0; j < (len(A[i])/2)+(len(A[i])%2); j++ {
			first := A[i][j]
			last := A[i][len(A[i])-j-1]
			A[i][j] = last^1
			A[i][len(A[i])-j-1] = first^1
		}
	}

	return A
}
