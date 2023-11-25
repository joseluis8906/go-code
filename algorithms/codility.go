package algorithms

import "math"

func Solution(A []int) int {
	n := len(A)
	N := 10000010
	p := make([]bool, N)

	maxele := math.MinInt

	for i := 0; i < n; i++ {
		if A[i] > 0 && A[i] <= n {
			p[A[i]] = true
		}

		maxele = int(math.Max(float64(maxele), float64(A[i])))
	}

	for i := 1; i < N; i++ {
		if !p[i] {
			return i
		}
	}

	return maxele + 1
}
