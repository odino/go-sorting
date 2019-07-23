package main

import (
	"math/rand"
)


func NewList(size int) ([]int) {
	l := make([]int, size, size)

	for i := 0; i < size; i++ {
		l[i] = rand.Intn(100)
	}

	return l
}

func main(){}