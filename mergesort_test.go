package main

import (
	"testing"
	"sort"
	"reflect"
)

func TestMergeSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		l := NewList(i)
		
		m := make([]int, i)
		s := make([]int, i)
		copy(m, l)
		copy(s, l)
		m = mergesort(m)
		sort.Ints(s)

		if !reflect.DeepEqual(m, s) {
			t.Errorf("%v is not %v", m, s)
		}
	}	
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewList(i)
		
		mergesort(l)
	}
}