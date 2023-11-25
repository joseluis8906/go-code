package algorithms

func QuickSort(a []int, p, r int) {
	if p < r {
		q := partition(a, p, r)
		QuickSort(a, p, q-1)
		QuickSort(a, q+1, r)
	}
}

func partition(a []int, p, r int) int {
	x := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] <= x {
			exchange(&a[i], &a[j])
			i++
		}
	}

	exchange(&a[i], &a[r])

	return i
}

func exchange(a, b *int) {
	y := *a
	*a = *b
	*b = y
}
