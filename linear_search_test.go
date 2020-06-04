package main

import "testing"

func prepareArray() [100]int {
	var a [100]int
	for i := 0; i < 100; i++ {
		a[i] = i*10 + 1
	}
	return a
}

func prepareTable() []struct {
	target   int
	expected int
} {
	return []struct {
		target   int
		expected int
	}{
		{0, NotFound},
		{10, NotFound},
		{1, 0},
		{11, 1},
		{101, 10},
		{991, 99},
		{992, NotFound},
	}
}

func TestLinearSearch(t *testing.T) {
	a := prepareArray()
	tests := prepareTable()

	for _, tt := range tests {
		actual := linearSearch(a, tt.target)
		if tt.expected != actual {
			t.Errorf("expected %d but got %d\n", tt.expected, actual)
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	a := prepareArray()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			linearSearch(a, j)
		}
	}
}
