# go-sorting

Various sorting algorithms implemented in Go.

```
$ make

go test -bench=. ./... -benchtime=1000x
goos: linux
goarch: amd64
BenchmarkBubbleSort-8   	    1000	    261670 ns/op
BenchmarkMergeSort-8    	    1000	     96989 ns/op
PASS
ok  	_/home/alex/projects/go-sorting	0.365s
```

## Why?

Why not?

## But, really, why?

For the ❤ of Go.

## List

* [mergesort](https://github.com/odino/go-sorting/blob/master/mergesort.go)
* [bubblesort](https://github.com/odino/go-sorting/blob/master/bubblesort.go)