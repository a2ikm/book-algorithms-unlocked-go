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
