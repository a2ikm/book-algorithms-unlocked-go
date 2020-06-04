package main

import "testing"

func TestLinearSearch(t *testing.T) {
	var a [100]int
	for i := 0; i < 100; i++ {
		a[i] = i*10 + 1
	}

	tests := []struct {
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

	for _, tt := range tests {
		actual := linearSearch(a, 100, tt.target)
		if tt.expected != actual {
			t.Errorf("expected %d but got %d\n", tt.expected, actual)
		}
	}
}
