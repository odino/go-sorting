package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func list(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

type algorithm func([]int) []int

func run(algo string, fn algorithm) {
	fmt.Printf("=== %s ===\n\n", strings.ToUpper(algo))
	for x := 0; x < 10; x++ {
		l := list(x)
		fmt.Printf("Original: %d\nSorted: %d\n\n", l, fn(l))
	}
}

func main() {
	run("mergesort", mergesort)
}
