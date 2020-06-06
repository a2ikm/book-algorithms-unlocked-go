package main

func BinarySearch(a [100]int, x int) (int, bool) {
	p := 0
	r := len(a) - 1
	for p <= r {
		q := (p + r) / 2
		if a[q] == x {
			return q, true
		} else if a[q] < x {
			p = q + 1
		} else {
			r = q - 1
		}
	}
	return 0, false
}

func RecursiveBinarySearch(a [100]int, x int) (int, bool) {
	p := 0
	r := len(a) - 1
	return internalRecursiveBinarySearch(a, p, r, x)
}

func internalRecursiveBinarySearch(a [100]int, p, r, x int) (int, bool) {
	if p > r {
		return 0, false
	}

	q := (p + r) / 2
	if a[q] == x {
		return q, true
	} else if a[q] < x {
		return internalRecursiveBinarySearch(a, q+1, r, x)
	} else {
		return internalRecursiveBinarySearch(a, p, r-1, x)
	}
}
