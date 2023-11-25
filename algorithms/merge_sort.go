package algorithms

import (
	"math"
)

func MergeSort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(a, p, q)
		MergeSort(a, q+1, r)
		merge(a, p, q, r)
	}
}

func merge(a []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q

	left := make([]int, n1+1)
	right := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		left[i] = a[p+i]
	}

	for j := 0; j < n2; j++ {
		right[j] = a[q+j+1]
	}

	left[n1] = math.MaxInt
	right[n2] = math.MaxInt

	var i, j int

	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
	}
}
