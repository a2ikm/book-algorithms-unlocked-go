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

func betterLinearSearch(a [100]int, x int) int {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == x {
			return i
		}
	}
	return NotFound
}
func sentinelLinearSearch(a [100]int, x int) int {
	n := len(a)
	last := a[n-1]
	a[n-1] = x
	i := 0
	for a[i] != x {
		i += 1
	}
	a[n-1] = last
	if a[i] == x {
		return i
	} else {
		return NotFound
	}
}
