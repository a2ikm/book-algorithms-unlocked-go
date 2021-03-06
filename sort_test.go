package main

import "testing"

func prepareArray3() []int {
	return []int{9, 1, 8, 2, 7, 3, 7, 4, 6, 5}
}

func TestSelectionSort(t *testing.T) {
	a := prepareArray3()
	b := SelectionSort(a)

	for i := 0; i < len(b)-1; i++ {
		if b[i] > b[i+1] {
			t.Errorf("expected b[i] <= b[i+1] but got %d > %d at i = %d\n", b[i], b[i+1], i)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := prepareArray3()
		SelectionSort(a)
	}
}

func TestMergeSort(t *testing.T) {
	a := prepareArray3()
	MergeSort(a)

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Errorf("expected b[i] <= b[i+1] but got %d > %d at i = %d\n", a[i], a[i+1], i)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := prepareArray3()
		MergeSort(a)
	}
}

func TestQuickSort(t *testing.T) {
	a := prepareArray3()
	QuickSort(a)

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			t.Errorf("expected b[i] <= b[i+1] but got %d > %d at i = %d\n", a[i], a[i+1], i)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := prepareArray3()
		QuickSort(a)
	}
}
