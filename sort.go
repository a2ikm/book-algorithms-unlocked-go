package main

func SelectionSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		s := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[s] {
				s = j
			}
		}
		a[i], a[s] = a[s], a[i]
	}
	return a
}

func MergeSort(a []int) {
	mergeSort(a, 0, len(a)-1)
}

func mergeSort(a []int, p, r int) {
	if p >= r {
		return
	}

	q := (p + r) / 2
	mergeSort(a, p, q)   // partial array: B
	mergeSort(a, q+1, r) // partial array: C
	merge(a, p, q, r)
}

func merge(a []int, p, q, r int) {
	i := p     // partial array: B
	j := q + 1 // partial array: C
	for k := p; k <= r; k++ {
		if i <= q && j <= r {
			if a[i] < a[j] {
				a[k] = a[i]
				i += 1
			} else {
				a[k] = a[j]
				j += 1
			}
		} else if i <= q {
			for i <= q {
				a[k] = a[i]
				i += 1
				k += 1
			}
		} else if j <= r {
			for j <= r {
				a[k] = a[j]
				j += 1
				k += 1
			}
		}
	}
}
