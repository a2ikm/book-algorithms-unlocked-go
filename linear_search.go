package main

var NotFound = -1

func linearSearch(a [100]int, x int) int {
	ret := NotFound
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == x {
			ret = i
		}
	}
	return ret
}
