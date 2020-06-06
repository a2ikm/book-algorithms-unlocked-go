package main

import "testing"

func prepareArray2() [100]int {
	var a [100]int
	for i := 0; i < 100; i++ {
		a[i] = i*10 + 1
	}
	return a
}

func prepareTable2() []struct {
	target int
	index  int
	found  bool
} {
	return []struct {
		target int
		index  int
		found  bool
	}{
		{0, 0, false},
		{10, 0, false},
		{1, 0, true},
		{11, 1, true},
		{101, 10, true},
		{991, 99, true},
		{992, 0, false},
	}
}

func TestBinarySearch(t *testing.T) {
	a := prepareArray2()
	tests := prepareTable2()

	for _, tt := range tests {
		i, found := BinarySearch(a, tt.target)
		if tt.found {
			if tt.found != found || tt.index != i {
				t.Errorf("expected (%d, %t) but got (%d, %t)\n", tt.index, tt.found, i, found)
			}
		} else {
			if tt.found != found {
				t.Errorf("expected (_, %t) but got (_, %t)\n", tt.found, found)
			}
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	a := prepareArray2()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			BinarySearch(a, j)
		}
	}
}
