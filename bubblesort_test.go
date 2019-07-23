package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for i := 10; i < 11; i++ {
		l := NewList(i)

		m := make([]int, i)
		s := make([]int, i)
		copy(m, l)
		copy(s, l)
		m = bubblesort(m)
		sort.Ints(s)

		if !reflect.DeepEqual(m, s) {
			t.Errorf("%v is not %v", m, s)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewList(i)

		bubblesort(l)
	}
}
